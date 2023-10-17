# Bare minimum HTTP echo service by Go net/http

Similar to the well-known echo service defined in RFC862, but works on HTTP/1.1.

Once a HTTP connection is established, any data received as request body is sent back as response body.

This works as full-duplex streaming thanks to HTTP/1.1 chunked encoding. However, note that not every HTTP/1.1 client can handle this kind of full-duplex HTTP streaming well.

In Go, we cannot just io.Copy() to http.ResponseWriter due to it's buffering. We need to call http.flusher.Flush() on every write to send immediate echo back to the client. 