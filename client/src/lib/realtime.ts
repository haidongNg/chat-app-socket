import ioClient from 'socket.io-client';
const ENDPOINT = "http://localhost:8080";

const socket = ioClient(ENDPOINT);

export const socketClient = socket;