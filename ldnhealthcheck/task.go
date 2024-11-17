package ldnhealthcheck

import (
	"RyuLdnWebsite/ldnhealthcheck/packets"
	"sync/atomic"
	"time"
)

type RyuLdnTests struct {
	PingTest   bool      `json:"ping_test"`
	InitTest   bool      `json:"init_test"`
	APTest     bool      `json:"ap_test"`
	UpdateTime time.Time `json:"update_time"`
}

var ryuLdnTestsPtr atomic.Pointer[RyuLdnTests]

func initializeTests() {
	initialTests := &RyuLdnTests{
		PingTest: false,
		InitTest: false,
		APTest:   false,
	}
	ryuLdnTestsPtr.Store(initialTests)
}

func updateTests(newTests RyuLdnTests) {
	ryuLdnTestsPtr.Store(&newTests)
}

func GetTests() *RyuLdnTests {
	return ryuLdnTestsPtr.Load()
}

func test(ldnHost string, ldnPort int, timeout time.Duration) {
	pingTest := false
	initTest := false
	apTest := false

	client, err := NewRyujinxLdnClient(ldnHost, ldnPort, timeout)
	if err == nil {
		pingTest = true
		updateTests(RyuLdnTests{PingTest: pingTest, InitTest: initTest, APTest: apTest, UpdateTime: time.Now()})
	} else {
		return
	}

	defer client.Close()

	initPacket := packets.NewInitializePacket()
	if err := client.Send(initPacket); err == nil {
		initTest = true
		updateTests(RyuLdnTests{PingTest: pingTest, InitTest: initTest, APTest: apTest, UpdateTime: time.Now()})
	} else {
		return
	}

	_, err = client.Receive()
	if err == nil {
		apTest = true
		updateTests(RyuLdnTests{PingTest: pingTest, InitTest: initTest, APTest: apTest, UpdateTime: time.Now()})
	} else {
		return
	}
}

func SceduleTask(repeat time.Duration, ldnHost string, ldnPort int, timeout time.Duration) {
	initializeTests()

	go func() {
		ticker := time.NewTicker(repeat)
		defer ticker.Stop()

		test(ldnHost, ldnPort, timeout)
		for range ticker.C {
			test(ldnHost, ldnPort, timeout)
		}
	}()
}
