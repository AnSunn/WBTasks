# How to run the program:

1. you can specify other host + port or another timeout
```
go run task.go --timeout=100s stackoverflow.com 80 
```
2. Write the server request in stdin, following this example 

**NB**: the subsequent lines are mandatory for a GET request, and end with two blank lines
```
GET /questions HTTP/1.0
Host: stackoverflow.com


```

3. To conclude the request enter '^]'
```
^]
```
So your request has to look like:

![request.png](..%2F..%2F..%2F..%2F..%2FDesktop%2Frequest.png)

4. Press enter

You receive the answer from the server. 

**NB** if the server sends 403 forbidden, then it closes the connection and next time sending the request you receive the info

For the example provided above, the server responds with:
```
HTTP/1.1 403 Forbidden
Date: Tue, 16 Jan 2024 10:06:40 GMT
Content-Type: text/html; charset=UTF-8
Content-Length: 4519
Connection: close
...
<head>
<title>Attention Required! | Cloudflare</title>
<meta charset="UTF-8" />
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
...
```
