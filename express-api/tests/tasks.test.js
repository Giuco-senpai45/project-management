const request = require("supertest");
const createServer = require("../utils/server");

const getTasksMock = jest.fn();
const createTaskMock = jest.fn();
const getTaskByIdMock = jest.fn();
const updateTaskMock = jest.fn();

const app = createServer({
  getTasks: getTasksMock,
  getTaskById: getTaskByIdMock,
  createTask: createTaskMock,
  updateTask: updateTaskMock,
});
let createdTaskId;

describe("GET /tasks", () => {
  it("should return 200 and a list of existing tasks (200 list of existing tasks)", async () => {
    const res = await request(app).get("/api/v1/tasks");
    expect(res.statusCode).toEqual(200);
    return expect(res.type).toEqual("application/json");
  });
});

describe("POST /tasks", () => {
  it("creates a new task (201 created task)", async () => {
    const res = await request(app).post("/api/v1/tasks").send({
      name: "TestTask",
      description: "This is the description of the task",
      status: false,
      urgency: 0,
      project_id: 1,
    });
    expect(res.statusCode).toEqual(201);
    expect(res.body).toHaveProperty("id");
    createdTaskId = res.body.id;
    expect(res.body.name).toEqual("TestTask");
    expect(res.body.description).toEqual("This is the description of the task");
    expect(res.body.status).toEqual(false);
    return expect(res.body.urgency).toEqual("0");
  });
});

describe("GET /tasks/:id", () => {
  it("find an existing task by id (200 ok, found task)", async () => {
    const res = await request(app).get(`/api/v1/tasks/${createdTaskId}`);
    expect(res.body.name).toEqual("TestTask");
    expect(res.body.description).toEqual("This is the description of the task");
    expect(res.body.status).toEqual(false);
    return expect(res.body.urgency).toEqual("0");
  });

  it("find a non existing task (404 task not found)", async () => {
    const res = await request(app).get(`/api/v1/tasks/0`);
    expect(res.statusCode).toEqual(404);
    expect(res.body).toHaveProperty("message");
    return expect(res.body.message).toEqual("Task not found");
  });
});

describe("PUT /tasks/:id", () => {
  it("updates an existing tasks data (200 ok updated task)", async () => {
    const res = await request(app).put(`/api/v1/tasks/${createdTaskId}`).send({
      name: "NewTaskName",
      description: "New task description",
      status: true,
      urgency: 1,
    });
    expect(res.statusCode).toEqual(200);
    expect(res.body.name).toEqual("NewTaskName");
    expect(res.body.description).toEqual("New task description");
    expect(res.body.status).toEqual(true);
    return expect(res.body.urgency).toEqual("1");
  });

  it("updates a non existing user (404 not found)", async () => {
    const res = await request(app).put("/api/v1/tasks/0").send({
      name: "NewTaskName",
      description: "New task description",
      status: true,
      urgency: 1,
    });
    expect(res.statusCode).toEqual(404);
    expect(res.body).toHaveProperty("message");
    return expect(res.body.message).toEqual("Task not found");
  });
});
