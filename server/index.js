const path = require("path") // built-in nodejs
const http = require("http"); // built-in nodejs
const express = require('express');
const socketIO = require('socket.io');
const { generateTextMessage, generateLocationMessage } = require('../template/message');

const app = express(); // http.createServer
const server = http.createServer(app);
const io = socketIO(server);

// serve static
const publicPath = path.join(__dirname + "/../public");
app.use(express.static(publicPath))

// chat app
io.on("connection", (socket) => {
  console.log("New user connect to server")

  // publish topic (message/email/custom event)
  socket.emit("welcome", generateTextMessage("admin", "Welcome to the chat app"))

  socket.broadcast.emit("newUser", generateTextMessage("admin", "New user joined"))

  // get message from user and emit
  socket.on("createMsg", (msg) => {
    io.emit("sendMsg", generateTextMessage(msg.from, msg.text))
  })

  // get location from user and emit
  socket.on("createLocation", (msg) => {
    io.emit("sendLocation", generateLocationMessage(msg.from, msg.latitude, msg.longitude))
  })

  socket.on("disconnect", () => {
    console.log("User disconnected")
  })
})

// port 
const port = process.env.PORT || 5000;
server.listen(port, () => {
  console.log(`Server running on port ${port}`)
})