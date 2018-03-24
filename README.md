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

1. Make a copy *routes.json.example* to *routes.json*
```bash
// Fot unix
cp routes.json.example routes.json
```

2. Build an image:
```bash
docker-compose build
```

3. Run a container:
```bash
docker-compose up
```

By default docker exposes port *8888*

Example:
```bash
curl -XGET http://localhost:8888/users

curl -X POST http://localhost:8888/register \
     -H 'content-type: application/json' \
     -d '{
	    "name":"john"
     }'
```