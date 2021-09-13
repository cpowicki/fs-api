# fs-api
A rest API for remotely manipulating a filesystem. 

### Getting Started

You can run `fs-api` with the included `run.sh` script. 

```
bash scripts/run.sh
```

You can also use [`make`]([https://www.gnu.org/software/make/) to:

* `linux` compiles for linux
* `darwin` compiles for darwin
* `dockerize` builds a local docker image
* `clean` clear local builds 
* `test` runs unit tests

To run the containerized app:

```
make dockerize
docker run -p 3030:3030 fs-api:latest [<ARGS>]
```

### CLI

```
./fsapi -h

Usage of ./fs-api:
  -port int
        the server port (default 3030)
  -root string
        the root directory to expose for browsing (default "$HOME")
```

### HTTP API

`GET /` 

Returns a listing of the metadata of objects in the configured root directory

`GET /:path`

Returns either:
* The file contents as string, if `path` is a file
* A listing of metadata objects if `path` is a directory