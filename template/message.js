const generateTextMessage = (from, text) => {
  return {
    from,
    text,
    reactedAt: Date.now()
  };
};

const generateLocationMessage = (from, latitude, longitude) => {
  return {
    from,
    latitude,
    longitude
  };
};

module.exports = {
  generateTextMessage,
  generateLocationMessage,
};
