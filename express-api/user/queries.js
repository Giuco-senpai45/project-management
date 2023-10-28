const getUsers = "SELECT * FROM users"
const getUserById =  "SELECT * FROM users WHERE id = $1"
const getUserByEmail = "SELECT * FROM users WHERE email = $1"
const createUser = "INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING *"
const updateUser = "UPDATE users SET name = $1, email = $2, password = $3, updated_at = $4 WHERE id = $5 RETURNING *"
const getUserProjects = "SELECT * FROM projects INNER JOIN user_projects ON projects.id = user_projects.project_id WHERE user_projects.user_id = $1"

module.exports = {
    getUsers,
    getUserById,
    getUserByEmail,
    createUser,
    updateUser,
    getUserProjects,
}