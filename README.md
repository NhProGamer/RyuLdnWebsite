# RyuLdnWebsite

RyuLdnWebsite is a Golang webserver compatible with RyuLdn backend.

## Features

- **Customizable Network Configuration**: Configurable via environment variables to meet various networking needs.
- **Proxy Support**: Enables you to specify proxy addresses in CIDR notation for secure and controlled access.

## Installation

Clone this repository:

```bash
git clone https://github.com/yourusername/RyuLdnWebsite.git
cd RyuLdnWebsite
```

Make sure you have Go installed, then build the project:

```bash
go build
# Don't forget to have public and static folders on the same directory of the binary
```

## Environment Variables

To configure RyuLdnWebsite, you can set the following environment variables:

| Variable              | Description                                                | Default Value | Example                                   |
|-----------------------|------------------------------------------------------------|---------------|-------------------------------------------|
| `LDN_HOST`            | IP or domain name of the RyuLdn server                     | -             | `192.168.1.100` or `ryu-ldn-server.local` |
| `LDN_PORT`            | Port of the RyuLdn server                                  | -             | `30456`                                    |
| `LDN_HEALTHCHECK_TIME`| Time between health checks in seconds                      | 300           | `300`                                     |
| `REDIS_URL`           | URL of the Redis server used for connection data storage   | -             | `192.168.1.152:5654`                      |
| `PROXIES`             | Comma-separated list of allowed IP ranges in CIDR notation | -             | `192.168.1.0/8,172.17.0.0/16`             |
| `HOST`                | Host address for the server                                | `0.0.0.0`     | `0.0.0.0`                                 |
| `PORT`                | Port for the server to listen on                           | `8080`        | `8080`                                    |


## Usage

To run the server with your custom configuration, make sure the environment variables are set, then start the server:

```bash
LDN_HOST=127.0.0.1 LDN_PORT=30456 LDN_HEALTHCHECK_TIME=300 REDIS_URL=192.168.1.152:5654 PROXIES=192.168.1.0/8,172.17.0.0/16 HOST=0.0.0.0 PORT=8080 ./RyuLdnWebsite
```

Or, to set variables in a `.env` file:

```bash
LDN_HOST=127.0.0.1
LDN_PORT=30456
LDN_HEALTHCHECK_TIME=300
REDIS_URL=192.168.1.152:5654
PROXIES=192.168.1.0/8,172.17.0.0/16
HOST=0.0.0.0
PORT=8080
```

Then start the server:

```bash
source .env
./RyuLdnWebsite
```
