const http = require('http');
const hostname = '127.0.0.1';
const port = 3000;

var server = http.createServer();
server.on('request', doRequest);

// ファイルモジュールを読み込む
var fs = require('fs');

var {PythonShell} = require('python-shell');
var options = {
    mode: 'text', // testもしくはjson
    pythonPath: '/Users/noriakioshita/.pyenv/shims/python3', 
    pythonOptions: ['-u'], 
  };

// リクエストの処理
function doRequest(req, res) {
    // ファイルを読み込んだら、コールバック関数を実行する。
    fs.readFile('index.html', 'utf-8' , doReard );
    
    // コンテンツを表示する。
    function doReard(err, data) {
        res.writeHead(200, {'Content-Type': 'text/html'});
        res.write(data);
        res.end();
    }
    
}

var pyshell = new PythonShell('make_kinmuhyo.py',options);

pyshell.on('message',function(message){
    console.log(message);
});

server.listen(port, hostname, () => {
    console.log(`Server running at http://${hostname}:${port}/`);
  });