var socket = io();

socket.on("connect", () => {
  console.log("Connect to server");
});

socket.on("welcome", msg => {
  console.log(msg);
});

socket.on("newUser", msg => {
  console.log(msg);
});

socket.on("sendMsg", msg => {
  console.log(msg);
  let liTag = $(
    `<li>${moment(msg.createdAt).format("h:mm a")} - ${msg.text}</li>`
  );
  $("#messages").append(liTag);
});

socket.on("disconnect", () => {
  console.log("Disconnect to server");
});

$("#message-form").on("submit", e => {
  e.preventDefault();

  socket.emit("createMsg", {
    from: "User",
    text: $("[name=message]").val()
  });
});

$("#sendLocation").on("click", () => {
  if (!navigator.geolocation)
    return alert("Your browser does not support this feature, please update");

  navigator.geolocation.getCurrentPosition(position => {
    // console.log(position.coords)
    const { latitude, longitude } = position.coords;

    socket.emit("createLocation", {
      from: "User",
      latitude,
      longitude
    });
  });
});

socket.on("sendLocation", msg => {
  let liTag = $("<li></li>");
  let aTag = $("<a>My location</a>");
  aTag.attr(
    "href",
    `https://www.google.com/maps/@${msg.latitude},${msg.longitude}z`
  );
  aTag.attr("target", "_blank");

  liTag.append(aTag);
  $("#messages").append(liTag);
});
