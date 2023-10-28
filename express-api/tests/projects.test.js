const request = require("supertest");
const createServer = require("../utils/server");

const getProjectsMock = jest.fn();
const createProjectMock = jest.fn();
const getProjectByIdMock = jest.fn();
const updateProjectMock = jest.fn();
const linkProjectToUserMock = jest.fn();

const app = createServer({
  getProjects: getProjectsMock,
  getUserById: getProjectByIdMock,
  createUser: createProjectMock,
  updateUser: updateProjectMock,
  linkProjectToUser: linkProjectToUserMock,
});

let createdProjectId;

// jest.mock("../helpers/sessions", () => ({
//   verifyToken: (_, __, next) => next()
// }));

describe("GET /projects", () => {
  it("should the list of existing projects", async () => {
    const res = await request(app).get("/api/v1/projects");
    expect(res.statusCode).toEqual(200);
    expect(res.type).toEqual("application/json");
  });
});

describe("POST /projects", () => {
  it("tries to create a new project without the user being logged in", async () => {
    const res = await request(app).post("/api/v1/projects").send({
      name: "TestProject",
      description: "The test projects description",
      urgency: 1,
      state: false,
    });
    expect(res.statusCode).toEqual(401);
  });
});
