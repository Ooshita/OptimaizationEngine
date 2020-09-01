function DockerExec() {
    var Docker = require('dockerode');
    var docker = new Docker();
    var elem = document.getElementById("output");

    docker.run('timesheets_image', ['bash', '-c', 'uname -a'], process.stdout, function (err, data ,container) {
        var output = data[0];
        console.log(data.StatusCode);
        elem.innerHTML = output;
    });
};