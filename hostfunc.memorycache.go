package capsule

import "github.com/valyala/fastjson"

/* Documentation
## Memory cache functions

### CacheSet

```golang
buffValue := capsule.CacheSet(key string, value []byte)
```

### CacheGet

```golang
buffValue, err := capsule.CacheGet(key string)
```

> if the err is not nil, you should get: "key not found"

*/

//export hostCacheSet
func hostCacheSet(
	keyPosition, keyLength uint32,
	valueStringPosition, valueStringLength uint32,
	returnValuePosition **uint32, returnValueLength *uint32) uint32
	
// CacheSet is an helper to use the hostCacheSet function
func CacheSet(key string, value []byte) []byte {

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
	data, _ := Result(bufferResponseFromHost)

	return data
}


//export hostCacheGet
func hostCacheGet(
	keyPosition, keyLength uint32,
	returnValuePosition **uint32, returnValueLength *uint32) uint32
	
// CacheGet is an helper to use the hostCacheGet function
func CacheGet(key string) ([]byte, error) {

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



//export hostCacheDel
func hostCacheDel(
	keyPosition, keyLength uint32,
	returnValuePosition **uint32, returnValueLength *uint32) uint32
	
// CacheDel is an helper to use the hostCacheDel function
func CacheDel(key string) []byte {

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
	data, _ := Result(bufferResponseFromHost)
	return data
}



//export hostCacheKeys
func hostCacheKeys(
	filterPosition, filterLength uint32,
	returnValuePosition **uint32, returnValueLength *uint32) uint32
	
// CacheKeys is an helper to use the hostCacheKeys function
func CacheKeys(filter string) ([]string, error) {

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
