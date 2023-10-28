const { Router } = require('express');
const controller = require('./controller');
const verifyToken = require('../helpers/sessions');

const router = Router();

router.get('/', controller.getProjects);
router.post('/', verifyToken, controller.createProject);
router.get('/:id', controller.getProjectById);
router.put('/:id', controller.updateProject);

module.exports = router;