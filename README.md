# docker-webdav
WebDAV and web page directory listing server

Default read/write and no authorization. I use it behind a proxy that
handles security.

## Usage

```sh
docker run --rm -v /path:/webdav -p 8080:8080 mwader/webdav
```

## License
Public domain
