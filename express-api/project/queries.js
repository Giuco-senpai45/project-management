const getProjects = "SELECT * FROM projects"
const getProjectById =  "SELECT * FROM projects WHERE id = $1"
const createProject = "INSERT INTO projects (name, description, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING *"
const updateProject = "UPDATE projects SET name = $1, description = $2, updated_at = $3 WHERE id = $4 RETURNING *"
const linkProjectToUser = "INSERT INTO user_projects (user_id, project_id) VALUES ($1, $2)"

module.exports = {
    getProjects,
    getProjectById,
    createProject,
    linkProjectToUser,
    updateProject,
}