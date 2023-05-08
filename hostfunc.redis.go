package capsule

import "github.com/valyala/fastjson"

// hostRedisSet sets the value of a redis key.
//
// keyPosition: the position of the key in memory.
// keyLength: the length of the key.
// valueStringPosition: the position of the value string in memory.
// valueStringLength: the length of the value string.
// returnValuePosition: the position of the pointer to the return value in memory.
// returnValueLength: the length of the return value.
//
// Returns the new uint32 value.
//export hostRedisSet
func hostRedisSet(
	keyPosition, keyLength uint32,
	valueStringPosition, valueStringLength uint32,
	returnValuePosition **uint32, returnValueLength *uint32) uint32


// RedisSet sends a message to the host to cache a key and its value.
//
// key: the string value of the key to cache.
// value: the byte slice value to cache.
// []byte: a byte slice that contains the response from the host.
func RedisSet(key string, value []byte) ([]byte, error) {

	keyPosition, keyLength := getBufferPosSize([]byte(key))
	valueStringPosition, valueStringLength := getBufferPosSize(value)

	// This will be use to get the response from the host
	var responseBufferPtr *uint32
	var responseBufferSize uint32

	// Send the lessage to the host
	hostCacheSet(
		keyPosition, keyLength,
		valueStringPosition, valueStringLength,
		&responseBufferPtr, &responseBufferSize)

	bufferResponseFromHost := readBufferFromMemory(responseBufferPtr, responseBufferSize)
	
	// check if success or failure
	data, err := Result(bufferResponseFromHost)

	return data, err
}



// hostRedisGet retrieves a Redis key-value pair from the host.
//
// keyPosition: the position of the key in memory.
// keyLength: the length of the key in memory.
// returnValuePosition: a pointer to the position of the return value in memory.
// returnValueLength: a pointer to the length of the return value in memory.
//
// Returns the new function.
//export hostRedisGet
func hostRedisGet(
	keyPosition, keyLength uint32,
	returnValuePosition **uint32, returnValueLength *uint32) uint32

// RedisGet retrieves the value for the given key from Redis.
//
// It takes a single string parameter, `key`, which is used to identify the value
// to retrieve from Redis.
//
// RedisGet returns a slice of bytes containing the retrieved value, and an error
// if the retrieval failed or the key was not found.
func RedisGet(key string) ([]byte, error) {

	keyPosition, keyLength := getBufferPosSize([]byte(key))

	// This will be use to get the response from the host
	var responseBufferPtr *uint32
	var responseBufferSize uint32

	// Send the lessage to the host
	hostCacheGet(
		keyPosition, keyLength,
		&responseBufferPtr, &responseBufferSize)

	bufferResponseFromHost := readBufferFromMemory(responseBufferPtr, responseBufferSize)
	
	// check if success or failure
	data, err := Result(bufferResponseFromHost)
	if err != nil {
		return nil, err // "key not found"
	}
	return data, nil
}




// hostRedisDel deletes the Redis key stored at the specified position and returns the length of the deleted key.
//
// keyPosition: the position of the key in Redis.
// keyLength: the length of the key.
// returnValuePosition: a pointer to the position of the return value.
// returnValueLength: a pointer to the length of the return value.
//
// Returns: the length of the deleted key.
//export hostRedisDel
func hostRedisDel(
	keyPosition, keyLength uint32,
	returnValuePosition **uint32, returnValueLength *uint32) uint32

// RedisDel deletes a Redis key and returns the result as a slice of bytes.
// The key parameter is a string representing the key to be deleted.
// The function returns a slice of bytes and an error.
func RedisDel(key string) ([]byte, error) {

	keyPosition, keyLength := getBufferPosSize([]byte(key))

	// This will be use to get the response from the host
	var responseBufferPtr *uint32
	var responseBufferSize uint32

	// Send the lessage to the host
	hostCacheDel(
		keyPosition, keyLength,
		&responseBufferPtr, &responseBufferSize)

	bufferResponseFromHost := readBufferFromMemory(responseBufferPtr, responseBufferSize)
	
	// check if success or failure
	data, err := Result(bufferResponseFromHost)
	return data, err
}

// hostRedisKeys returns a uint32 representing the new function.
//
// filterPosition: uint32 representing the filter position.
// filterLength: uint32 representing the filter length.
// returnValuePosition: pointer to uint32 representing the return value position.
// returnValueLength: pointer to uint32 representing the return value length.
//
// Returns a uint32 representing the new function.
//export hostRedisKeys
func hostRedisKeys(
	filterPosition, filterLength uint32,
	returnValuePosition **uint32, returnValueLength *uint32) uint32

// RedisKeys returns an array of Redis keys that match the given filter.
//
// filter: A string used to filter Redis keys.
// Returns an array of strings and an error.
func RedisKeys(filter string) ([]string, error) {

	filterPosition, filterLength := getBufferPosSize([]byte(filter))

	// This will be use to get the response from the host
	var responseBufferPtr *uint32
	var responseBufferSize uint32

	// Send the lessage to the host
	hostCacheKeys(
		filterPosition, filterLength,
		&responseBufferPtr, &responseBufferSize)

	bufferResponseFromHost := readBufferFromMemory(responseBufferPtr, responseBufferSize)
	//! 🤚 this is a json string (array json string: `["Hello", "World"]`)
	
	// check if success or failure
	data, err := Result(bufferResponseFromHost)
	if err != nil {
		return nil, err
	} 
	var jsonParser fastjson.Parser
	keysArray, err := jsonParser.Parse(string(data))
	if err != nil {
		return nil, err
	}
	var keys []string
	for _, key := range keysArray.GetArray("keys") {
		keys = append(keys, string(key.GetStringBytes())) 
		//! if it doesn't work, implement my own simple parser
	}
	return keys, nil

}
