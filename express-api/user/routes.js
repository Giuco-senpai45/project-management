const { Router } = require('express');
const controller = require('./controller');
const verifyToken = require('../helpers/sessions');

const router = Router();

router.get('/',  controller.getUsers);
router.post('/', controller.createUser);
router.get('/:id', controller.getUserById);
router.put('/:id', controller.updateUser);
router.post('/login', controller.loginUser);
router.get('/:id/projects', controller.getUserProjects);

module.exports = router;