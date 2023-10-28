function createServer(database) {
  const express = require("express");
  const cookieParser = require("cookie-parser");
  const userRoutes = require("../user/routes");
  const projectRoutes = require("../project/routes");
  const taskRoutes = require("../task/routes");

  const apiVersionString = "/api/v1";
  const app = express();
  port = process.env.APP_PORT || 3000;

  // middleware to get responses as json
  app.use(express.json());
  app.use(cookieParser());

  app.get("/", (req, res) => res.send("Hello World with Express"));

  // User Resource
  app.use(`${apiVersionString}/users`, userRoutes);

  // Project resource
  app.use(`${apiVersionString}/projects`, projectRoutes);

  // Task resource
  app.use(`${apiVersionString}/tasks`, taskRoutes);

  app.get('/clear-cookie', (_, res) => {
    res.clearCookie('token');
    res.send('Cookie cleared');
  });

  return app;
}

module.exports = createServer;
