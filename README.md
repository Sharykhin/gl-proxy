General proxy server:
=====================

Currently proxy server doesn't support **https** or **CONNECT** method.
Routes are configured in *routes.json* file by the following rule:  

"regexp pattern": "destination server"

Requirements:
------------

[docker](https://www.docker.com/)

Usage:
-----

1. Build an image:
```bash
docker-compose build
```