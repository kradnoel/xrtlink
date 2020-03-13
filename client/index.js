const axios = require('axios').default
const express = require('express')
const path = require('path')
const dirname = path.dir

const root = 'http://localhost:8080'
const app = express()
const port = '3000'

//app.engine('html', require('ejs').renderFile);

//app.set('views', __dirname + '/views')
//app.engine('html', require('ejs').renderFile);
app.engine('pug', require('pug').__express);
//app.set('view engine', 'html')

app.use(express.static(path.join(__dirname, 'public')))
app.set('appName', 'XRTLINK')

app.get('/', function(req, res) {
  res.render('index.pug', null, null)
})

app.get('/:uid', function(req, res) {
  var uid = req.params.uid
  axios.get(root + '/' + uid).then((r) => {
    var data = r.data
    if (data.error == false) {
      res.redirect(data.data)
    } else {
      //res.send('link doesnt exists')
      res.render('404.pug', null, null)
    }
  }).catch((e) => {
    console.log(e.errno)
    res.render('500.pug', null, null)
  })
})

/*app.get('/css/:file', function(req, res) {
  var file = req.params.file
  res.sendfile('public/css/' + file);
})

app.get('/js/:file', function(req, res) {
  var file = req.params.file
  res.sendfile('public/js/' + file);
})

app.get('/fonts/:file', function(req, res) {
  var file = req.params.file
  res.sendfile('public/fonts/' + file);
})

app.get('/img/:file', function(req, res) {
  var file = req.params.file
  res.sendfile('public/img/' + file);
})*/

app.listen(port, function() {
  console.log('The server is running, '+
    ' please, open your browser at http://localhost:%s', port)
})
