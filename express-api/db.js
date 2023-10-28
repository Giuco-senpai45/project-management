const Pool = require('pg').Pool;

const pool = new Pool({
    user: process.env.APP_PORT ||'postgres',
    host: process.env.HOST || 'localhost',
    database: process.env.DATABASE || 'projects-db',
    password: process.env.PASSWORD || 'password',
    port:  process.env.PORT || '5555',
});

module.exports = pool;