const pool = require("../db");
const queries = require("./queries");

const getProjects = (req, res) => {
  pool.query(queries.getProjects, (error, results) => {
    if (error) {
      throw error;
    }
    res.status(200).json(results.rows);
  });
};

const createProject = (req, res) => {
  const { name, description } = req.body;

  const loggedUser = req.user;

  // create project
  let createdProject;
  const createdAt = new Date();
  pool.query(
    queries.createProject,
    [name, description, createdAt.toISOString(), createdAt.toISOString()],
    (error, results) => {
      if (error) {
        throw error;
      }
      createdProject = results.rows[0];
      // insert project into composite table
      pool.query(
        queries.linkProjectToUser,
        [loggedUser.id, createdProject.id],
        (error, results) => {
          if (error) {
            throw error;
          }
        }
      );

      res.status(201).send(createdProject);
    }
  );
};

const getProjectById = (req, res) => {
  const id = parseInt(req.params.id);

  pool.query(queries.getProjectById, [id], (error, results) => {
    if (error) {
      throw error;
    }
    if (results.rows.length > 0) {
        res.status(200).json(results.rows[0]);
      } else {
        res.status(404).json({ message: "Project not found" });
      }
  });
};

const updateProject = (req, res) => {
  const id = parseInt(req.params.id);
  const { name, description } = req.body;

  const updatedAt = new Date();

  pool.query(
    queries.updateProject,
    [name, description, updatedAt.toISOString(), id],
    (error, results) => {
      if (error) {
        throw error;
      }

      if (results.rows.length === 0) {
        res.status(404).send({ message: "Project not found" });
      } else {
        res.status(200).send(results.rows[0]);
      }
    }
  );
};

module.exports = {
  getProjects,
  createProject,
  getProjectById,
  updateProject,
};
