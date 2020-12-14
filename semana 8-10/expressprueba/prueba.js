#!/usr/bin/env node

// Lo anterior es para poder usar systemd

var express=require('express');
var app = express();
var port = process.env.PORT || 8080;

app.get('/', function (req, res) {
    res.send( { Asigntaura: "CC",
                Opinion: "Estoy aprendiendo mucho",
                MeGusta: 2,
                NoMeGusta: 0} );
});

app.get('/helloworld', function (req, res) {
    res.send('Hello World !');
});

app.listen(port);
console.log('Server running at http://127.0.0.1:'+port+'/');

module.exports = app;