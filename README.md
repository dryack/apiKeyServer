### apiKeyServer
#### Serving Torn user keys for use by tools using the Torn API

Torn permits each _active_ user in game to make up to 100 requests per minute. For large tasks, it can be useful to move
beyond that limitation by using freely donated keys.  The problem I had was that each donated key had a different number
of uses/minute that were available since users want to be able to use their own tools/extensions as well.

apiKeyServer began as a project meant to address this problem, by allowing each key to have a specified number of 
uses/minute maximum, after which the key will not be available during a given minute.

##### Features
* gRPC interface using protobufs for messaging, keeps the server and communications lightweight, makes it easy to use 
TLS to secure the data (Torn requires us to protect user keys or face severe consequences).
* Many programs can all use the same server. gRPC can be used with C#, C++, Dart, Go, Java, Kotlin, Node, Objective-C, 
PHP, Python, and Ruby. I hear that JS is starting to receive support now as well.
* Keys can be individually configured for a maximum number of uses per minute
* Keys can be given one or more types, allowing clients to make specific requests, for example one could define a type 
'faction', which might be only those donated keys that have access to the author's faction's API endpoint.
* Leveled key usage: all else being equal, the key with the highest number of uses remaining will be returned to the 
client first.
* Because the server's minute may not really be synced to Torn's minute, sometimes the client will receive a Too Many 
Requests error code from the API. In this case the client Clients may tell the server to kill a key for the current 
minute.
* For a wide number of reasons, a key may become invalid (user has gone inactive or changed their key). If the client
receives an error code from the API that signals this, it may issue the server a command to permanently kill the key,
making it unavailable until the server is restarted (and the key is updated or removed from the configs). 


##### Planned features
- [ ] Options for logging; turn it off, change logfile location, etc.

- [ ] (Optionally) return a message signalling keys are exhausted, this would permit the client program to continue 
processing. Currently, we just wait until keys again become available before returning a result. This forces the client
to wait for a result or timeout. With an appropriate timeout, this has worked fine for me so far, but seems limiting as
a more general purpose solution.

- [ ] Request multiple keys in a single request. It may be more efficient for the client to request X keys of type Y at
once.