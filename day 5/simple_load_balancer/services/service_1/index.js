const Express = require('express');

const app = Express()

app.get('/', (req, res) => res.json("service 1"))

app.listen(8100, () => console.log("service 1 started at 127.0.0.1:8100"))
