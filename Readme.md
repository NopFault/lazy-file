# LazyFile

Lazy file - file server to send/receive file over `TCP`. File is "encoded" with `XOR` by the first file name letter to hide it in network flow.

### Server

To run it as a server:

    `lazyfile -h 0.0.0.0 -p 9999 -f filetosend`

To receive file:

    `lazyfile -h 0.0.0.0 -p 9999`
