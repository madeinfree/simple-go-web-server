# Simple Go Web Server

Run your small or test project server by golang.

# Features

- golang web server
- support .js file
- hot reload

# Uses

1. git clone https://github.com/madeinfree/simple-go-web-server
2. cd simple-go-web-server
3. alias simple-go-web-server='$(pwd)/bin/simple-go-web-server'
4. go to your website project directory
5. command ```simple-go-web-server [options]```
6. open http://localhost:3000

# Command Options

|    Options    | Type |          Description          | Default |
| -------------  | ---- |          -----------          | ----- |
|--port, -p|int|setting web server listen port| 3000 |
|--content, -c|string|setting hot loader file content base|null|
|--hot, -hh|boolean|should required when you use --content options| false |

# Port

When you would like to set your website port you can command

```command
simple-go-web-server -p 8080
```

to use 8080 port.

# Hot Reload

simple go web server help allow you to use hot reload, you can command

```command
$ simple-go-web-server -p 1234 -c "/build/bundle.js" --hot
```

and copy below script paste in your js file

```javascript
setInterval(() => {
  fetch('http://localhost:300/simple-go-server-file-change').then((r) => {
    r.json().then((rr) => {
      if (rr) {
        fetch('http://localhost:3000//simple-go-server-file-callback?isOK=true', {
          method: 'GET'
        })
        location.reload()
      }
    })
  })
}, 2000)
```
