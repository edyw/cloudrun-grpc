const express = require('express')
const { getContactHandler, getContactAuthHandler } = require('./contact-handler')

const app = express()
app.get('/contact/:phone', getContactHandler)
app.get('/contact-auth/:phone', getContactAuthHandler)
app.listen('8082', () => {
    console.log('node-api listening..')
})