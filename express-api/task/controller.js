const pool = require("../db");
const queries = require("./queries");

const getTasks = (req, res) => {
  pool.query(queries.getTasks, (error, results) => {
    if (error) {
      throw error;
    }
    res.status(200).json(results.rows);
  });
};

const createTask = (req, res) => {
  const { name, description, urgency, status, project_id } = req.body;

  const createdAt = new Date();
  pool.query(
    queries.createTask,
    [
      name,
      description,
      urgency,
      status,
      project_id,
      createdAt.toISOString(),
      createdAt.toISOString(),
    ],
    (error, results) => {
      if (error) {
        throw error;
      }

      res.status(201).send(results.rows[0]);
    }
  );
};

const getTaskById = (req, res) => {
  const id = parseInt(req.params.id);

  pool.query(queries.getTaskById, [id], (error, results) => {
    if (error) {
      throw error;
    }
    if (results.rows.length === 0) {
      res.status(404).send({ message: "Task not found" });
    } else {
      res.status(200).send(results.rows[0]);
    }
  });
};

const updateTask = (req, res) => {
  const id = parseInt(req.params.id);
  const { name, description, urgency, status } = req.body;
  const updatedAt = new Date();

  pool.query(
    queries.updateTask,
    [name, description, urgency, status, updatedAt.toISOString(), id],
    (error, results) => {
      if (error) {
        throw error;
      }
      if (results.rows.length === 0) {
        res.status(404).send({ message: "Task not found" });
      } else {
        res.status(200).send(results.rows[0]);
      }
    }
  );
};

module.exports = {
  getTasks,
  createTask,
  getTaskById,
  updateTask,
};
