# v3: Extended

Queries to different services.
Parameters:
* q : query
* s : service (if not found, default service will be selected)

Available services:
* google : Google search (default)
* maps : Google maps search
* instagram : instragram tags search
* twitter : twitter tags search
* flickr: flickr public pictures search

Adding services is possible by adding a service name in the json config file and adding a search url (leaving out the search value)

```js
"services": {
        "google" : "https://google.com/search?q="
}
```

Listen port is set in config.json
```js
"port": 8080
```