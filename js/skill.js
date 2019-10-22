const { json } = require("micro");

module.exports = async (req, res) => {
  const { request, session, version } = await json(req);

  res.end(
    JSON.stringify({
      version,
      session,
      response: {
        text: request.original_utterance || "hello"
      }
    })
  );
};
