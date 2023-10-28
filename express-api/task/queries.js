const getTasks = "SELECT * FROM tasks"
const getTaskById =  "SELECT * FROM tasks WHERE id = $1"
const createTask = "INSERT INTO tasks (name, description, urgency, status, project_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *"
const updateTask = "UPDATE tasks SET name = $1, description = $2, urgency=$3, status=$4, updated_at = $5 WHERE id = $6 RETURNING *"

module.exports = {
    getTasks,
    getTaskById,
    createTask,
    updateTask,
}