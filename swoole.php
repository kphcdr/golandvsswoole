<?php
$server = new swoole_http_server("0.0.0.0", 9001);


$server->on('request', function ($request, $response) {
    $response->end("hello world");
});

$server->start();