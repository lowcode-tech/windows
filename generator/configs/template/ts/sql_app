import { merge } from 'config-plus';
import dotenv from 'dotenv';
import express, { json } from 'express';
import { allow, MiddlewareLogger } from 'express-ext';
import http from 'http';
import { createLogger } from 'logger-core';
import mysql from 'mysql';
import { PoolManager } from 'mysql-core';
import { config, env } from './config';
import { useContext } from './context';
import { route } from './route';

dotenv.config();
const conf = merge(config, process.env, env, process.env.ENV);

const app = express();
const logger = createLogger(conf.log);
const middleware = new MiddlewareLogger(logger.info, conf.middleware);
app.use(allow(conf.allow), json(), middleware.log);

const pool = mysql.createPool(conf.db);
const db = new PoolManager(pool);
const ctx = useContext(db, logger, middleware);
route(app, ctx);
http.createServer(app).listen(conf.port, () => {
  console.log('Start server at port ' + conf.port);
});
