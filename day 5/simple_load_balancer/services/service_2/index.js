const Express = require('express');

const app = Express()

app.get('/', (req, res) => res.json("service 2"))

app.listen(8200, () => console.log("service 2 started at 127.0.0.1:8200"))
