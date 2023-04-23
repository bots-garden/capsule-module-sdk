"use strict";

import * as fs from 'fs'
import { WASI } from 'wasi'

import Fastify from 'fastify'
const fastify = Fastify({
  logger: true
})
import process from "node:process"



const wasi = new WASI({args: [""]})
const importObject = { wasi_snapshot_preview1: wasi.wasiImport };

function copyToMemory(bytes, instance) {
    // The TinyGo `malloc` is automatically exported
    const ptr = instance.exports.malloc(bytes.length);
    const mem = new Uint8Array(
        instance.exports.memory.buffer, ptr, bytes.length
    );
    mem.set(bytes);

    return {
        position: ptr,
        size: bytes.length
    };
}

function callHandle(position, size, instance) {
    // Call the `helloWorld` TinyGo function
    // Get a kind of pair of values
    const pointerSize = instance.exports.callHandle(position, size);

    const memory = instance.exports.memory;
    const completeBufferFromMemory = new Uint8Array(memory.buffer);

    const MASK = (2n ** 32n) - 1n;

    // Extract the values of the pair (using the mask)
    const ptrPosition = Number(pointerSize >> BigInt(32));
    const stringSize = Number(pointerSize & MASK);

    // Extract the string from the memory buffer
    const extractedBuffer = completeBufferFromMemory.slice(
        ptrPosition, ptrPosition + stringSize
    );
    
    const result = {};

    if (extractedBuffer[0] == 83) { // => success
        result.isSuccess = true
        result.isFailure = false
    } else { // => failure
        result.isSuccess = false
        result.isFailure = true
    }
    result.buffer = extractedBuffer.slice(1);
    return result;

}

// Create and start the HTTP server
const opts = {}

const start = async () => {

    const wasm = await WebAssembly.compile(fs.readFileSync("./function/main.wasm"));

    const instance = await WebAssembly.instantiate(wasm, importObject);

    wasi.start(instance);

    fastify.post('/', opts, async (request, reply) => {
        const stringParameter = request.body;
        const bytes = new TextEncoder("utf8").encode(stringParameter);
    
        let {position, size} = copyToMemory(bytes, instance);
        
        let {buffer, isFailure, isSuccess} = callHandle(position, size, instance);
    
        if(isSuccess) {
            // Decode the buffer
            const str = new TextDecoder("utf8").decode(buffer);
            return str
        } else {
            const str = new TextDecoder("utf8").decode(buffer);
            return str
        }
    })
  
    try {
      await fastify.listen({ port: 8080, host: '0.0.0.0'})
    } catch (err) {
      fastify.log.error(err)
      process.exit(1)
    }
  }
  start().then(r => console.log("😄 started"))