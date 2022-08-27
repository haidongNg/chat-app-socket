import fastify from 'fastify';
import fastifyIO from '@fastify/websocket';
import cors from '@fastify/cors'

const server = fastify();
server.register(cors)
//
server.register(fastifyIO, { options: { maxPayload: 1048576 } });


server.get('/*', {websocket: true}, (connection /* SocketStream */, req /* FastifyRequest */) => {
  console.log(req.id)
  connection.socket.on('message', message => {
    console.log(message)
    connection.socket.send({
      from: 'client',
      message: 'test socket fastify',
      time: new Date().toLocaleString(),
    });
  });
});

server.listen({ port: 8080 }, (err, address) => {
  if (err) {
    console.error(err)
    process.exit(1)
  }
  console.log(`Server listening at ${address}`)
})