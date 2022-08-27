import { FastifyPlugin } from 'fastify';
interface PluginOptions {
    //...
}

declare module 'fastify' {
    interface FastifyInstance {
        //...
    }
}