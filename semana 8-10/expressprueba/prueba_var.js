var express=require('express');
var app = express();
var port = process.env.PORT || 8080;

var array_resenias = [ { Asigntaura: "CC",
                        Opinion: "Estoy aprendiendo mucho",
                        MeGusta: 2,
                        NoMeGusta: 0},

                        { Asigntaura: "AA",
                        Opinion: "No me gusta nada",
                        MeGusta: 2,
                        NoMeGusta: 3},

                        { Asigntaura: "BB",
                        Opinion: "No está mal",
                        MeGusta: 0,
                        NoMeGusta: 2}
                    ]

app.get('/resenia', function (req, res) {
    res.send(array_resenias);
});

app.get('/resenia/:id', function (req, res) {
    if (req.params.id < array_resenias.length)
        res.send(array_resenias[req.params.id]);
    else    
        res.send("No hay tantas reseñas registradas");
});


app.get('/helloworld', function (req, res) {
    res.send('Hello World !');
});

app.listen(port);
console.log('Server running at http://127.0.0.1:'+port+'/');