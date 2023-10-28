const pool = require("../db");
const queries = require("./queries");
const hashing = require("../helpers/hashing");
const jwt = require("jsonwebtoken");
// const cookieParser = require('cookie-parser');

const getUsers = (_, res) => {
  pool.query(queries.getUsers, (error, results) => {
    if (error) {
      res.status(500).json({ error: error.message });
      return;
    }
    res.status(200).json(results.rows);
  });
};

const getUserById = (req, res) => {
  const id = parseInt(req.params.id);

  pool.query(queries.getUserById, [id], (error, results) => {
    if (error) {
      res.status(500).json({ error: error.message });
      return;
    }
    if (results.rows.length > 0) {
      res.status(200).json(results.rows[0]);
    } else {
      res.status(404).json({ message: "User not found" });
    }
  });
};

const createUser = (req, res) => {
  const { name, email, password } = req.body;

  hashing.cryptPassword(password, (err, hash) => {
    if (err) {
      console.error("Error encrypting password:", err);
      res.status(500).json({ error: error.message });
      return;
    } else {
      const createdAt = new Date();
      pool.query(
        queries.createUser,
        [name, email, hash, createdAt.toISOString(), createdAt.toISOString()],
        (error, results) => {
          if (error) {
            res.status(500).json({ error: error.message });
            return;
          }
          res.status(201).send(results.rows[0]);
        }
      );
    }
  });
};

const updateUser = (req, res) => {
  const id = parseInt(req.params.id);
  const { name, email, password } = req.body;

  hashing.cryptPassword(password, (err, hash) => {
    if (err) {
      console.error("Error encrypting password:", err);
      res.status(500).json({ error: error.message });
      return;
    } else {
      const updatedAt = new Date();
      pool.query(
        queries.updateUser,
        [name, email, hash, updatedAt.toISOString(), id],
        (error, results) => {
          if (error) {
            res.status(500).json({ error: error.message });
            return;
          }
          if (results.rows.length === 0) {
            res.status(404).send({ message: "User not found" });
          } else {
            res.status(200).send(results.rows[0]);
          }
        }
      );
    }
  });
};

const loginUser = (req, res) => {
  const { email, password } = req.body;

  let user;
  pool.query(queries.getUserByEmail, [email], (error, results) => {
    if (error) {
      res.status(500).json({ error: error.message });
      return;
    }

    if (results.rows.length > 0) {
      user = results.rows[0];
      hashing.comparePassword(
        password,
        user.password,
        (_, isPasswordMatch) => {
          if (!isPasswordMatch) {
            res.status(401).json({ message: "Invalid credentials" });
            return;
          } else {
            const token = jwt.sign({ id: user.id }, "secret_key");

            res.cookie("token", token, { httpOnly: true });

            res.status(200).json({ message: "Logged in successfully" });
          }
        }
      );
    } else {
      res.status(404).json({ message: "User not found" });
    }
  });
};

const getUserProjects = (req, res) => {
  const id = parseInt(req.params.id);

  pool.query(queries.getUserProjects, [id], (error, results) => {
    if (error) {
      res.status(500).json({ error: error.message });
      return;
    }
    res.status(200).json(results.rows);
  });
};

module.exports = {
  getUsers,
  getUserById,
  createUser,
  updateUser,
  loginUser,
  getUserProjects,
};
