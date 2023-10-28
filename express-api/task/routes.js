const { Router } = require('express');
const controller = require('./controller');
const verifyToken = require('../helpers/sessions');

const router = Router();

router.get('/', controller.getTasks);
router.post('/', controller.createTask);
router.get('/:id', controller.getTaskById);
router.put('/:id', controller.updateTask);

module.exports = router;