var request = require('supertest'),
app = require('../prueba.js');

request(app)
  .get('/')
  .expect('Content-Type', /json/)
  .expect(200)
  .end(function(err, res) {
    if (err) throw err;
  });


request(app)
  .get('/helloworld')
  .expect(200)
  .end(function(err, res) {
    if (err) throw err;
});

/* DESCOMENTAR PARA OBTENER UN ERROR
request(app)
  .get('/novale')
  .expect(200)
  .end(function(err, res) {
    if (err) throw err;
});
*/