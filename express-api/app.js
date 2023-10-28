const pool = require("./db");
const createServer = require("./utils/server");

const app = createServer(pool); 

app.listen(port);
console.log('API started on port: ' + port);