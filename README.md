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
don't forget to have public and static folders
```

## Environment Variables

To configure RyuLdnWebsite, you can set the following environment variables:

| Variable   | Description                                               | Default Value   | Example                           |
|------------|-----------------------------------------------------------|-----------------|-----------------------------------|
| `REDIS_URL` | URL of the Redis server used for connection data storage | -               | `192.168.1.152:5654`              |
| `PROXIES`   | Comma-separated list of allowed IP ranges in CIDR notation | -               | `192.168.1.0/8,172.17.0.0/16`     |
| `HOST`      | Host address for the server                              | `0.0.0.0`       | `0.0.0.0`                         |
| `PORT`      | Port for the server to listen on                         | `8080`          | `8080`                            |

## Usage

To run the server with your custom configuration, make sure the environment variables are set, then start the server:

```bash
REDIS_URL=192.168.1.152:5654 PROXIES=192.168.1.0/8,172.17.0.0/16 HOST=0.0.0.0 PORT=8080 ./RyuLdnWebsite
```

Or, to set variables in a `.env` file:

```bash
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
