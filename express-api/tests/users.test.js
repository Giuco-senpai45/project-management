const request = require("supertest");
const createServer = require("../utils/server");

const getUsersMock = jest.fn();
const createUserMock = jest.fn();
const getUserByIdMock = jest.fn();
const getUserByEmailMock = jest.fn();
const updateUserMock = jest.fn();
const getUserProjectsMock = jest.fn();

const app = createServer({
  getUsers: getUsersMock,
  getUserById: getUserByIdMock,
  getUserByEmail: getUserByEmailMock,
  createUser: createUserMock,
  updateUser: updateUserMock,
  getUserProjects: getUserProjectsMock,
});
let createdUserId;

jest.mock("jsonwebtoken", () => ({
  sign: jest.fn(() => "token_string"),
}));

describe("GET /users", () => {
  it("should return 200 and a list of users (200 list of existing users)", async () => {
    const res = await request(app).get("/api/v1/users");
    expect(res.statusCode).toEqual(200);
    return expect(res.type).toEqual("application/json");
  });
});

describe("POST /users", () => {
  it("creates a new user (201 created user)", async () => {
    const res = await request(app)
      .post("/api/v1/users")
      .send({
        name: "John Doe",
        email: "john.doe@example.com",
        password: "ps1",
      });
    expect(res.statusCode).toEqual(201);
    expect(res.body).toHaveProperty("id");
    createdUserId = res.body.id;
    expect(res.body.name).toEqual("John Doe");
    return expect(res.body.email).toEqual("john.doe@example.com");
  });
});

describe("GET /users/:id", () => {
  it("find an existing user by id (200 ok, found user)", async () => {
    const res = await request(app).get(`/api/v1/users/${createdUserId}`);
    expect(res.statusCode).toEqual(200);
    expect(res.body).toHaveProperty("id");
    expect(res.body.name).toEqual("John Doe");
    return expect(res.body.email).toEqual("john.doe@example.com");
  });

  it("find a non existing user (404 user not found)", async () => {
    const res = await request(app).get(`/api/v1/users/0`);
    expect(res.statusCode).toEqual(404);
    expect(res.body).toHaveProperty("message");
    return expect(res.body.message).toEqual("User not found");
  });
});

describe("POST /users/login", () => {
  it("logs in an existing user and generates a session storing a jwt in a cookie (200 ok,generates cookie)", async () => {
    const res = await request(app)
      .post("/api/v1/users/login")
      .send({ email: "john.doe@example.com", password: "ps1" });
    expect(res.statusCode).toEqual(200);
    expect(res.headers["set-cookie"]).toBeDefined();
    return expect(res.headers["set-cookie"][0]).toContain("token=token_string");
  });

  it("try login with non existing email (404 user not found)", async () => {
    const res = await request(app)
      .post("/api/v1/users/login")
      .send({ email: "nonexistent@example.com", password: "ps1" });
    expect(res.statusCode).toEqual(404);
    expect(res.body).toHaveProperty("message");
    return expect(res.body.message).toEqual("User not found");
  });
});

describe("POST /users/login", () => {
  it("logs in with wrong password (401 unauthorized)", async () => {
    const res = await request(app)
      .post("/api/v1/users/login")
      .send({ email: "john.doe@example.com", password: "wrongpas" });
    expect(res.statusCode).toEqual(401);
    expect(res.body).toHaveProperty("message");
    return expect(res.body.message).toEqual("Invalid credentials");
  });
});

describe("PUT /users/:id", () => {
  it("updates an existing users data (200 ok updated user)", async () => {
    const res = await request(app)
      .put(`/api/v1/users/${createdUserId}`)
      .send({
        name: "New Name",
        email: "newmail@example.com",
        password: "newpas",
      });
    expect(res.statusCode).toEqual(200);
    expect(res.body).toHaveProperty("id");
    expect(res.body.name).toEqual("New Name");
    return expect(res.body.email).toEqual("newmail@example.com");
  });

  it("updates a non existing user (404 not found)", async () => {
    const res = await request(app)
      .put("/api/v1/users/0")
      .send({
        name: "New Name",
        email: "newmail@example.com",
        password: "newpas",
      });
    expect(res.statusCode).toEqual(404);
    expect(res.body).toHaveProperty("message");
    return expect(res.body.message).toEqual("User not found");
  });
});

