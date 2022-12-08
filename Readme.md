# LazyFile

Lazy file - file server to send/receive file over `TCP`. File is "encoded" with `XOR` by the first file name letter to hide it in network flow.

### Server

To run it as a server:

    `lazyfile -h 0.0.0.0 -p 9999 -f filetosend`

To receive file:

    `lazyfile -h 0.0.0.0 -p 9999`

### POC

* Sending file data:

<img width="654" alt="Screenshot 2022-12-08 at 21 21 29" src="https://user-images.githubusercontent.com/90475186/206548317-e6bbcf01-e873-46ff-80d8-3d5a228af4ee.png">

* Start the server:

<img width="585" alt="Screenshot 2022-12-08 at 21 09 40" src="https://user-images.githubusercontent.com/90475186/206545943-1c97bd7c-3a8c-40a2-81b4-cc115bdd8e5b.png">

* Receive file:

<img width="622" alt="Screenshot 2022-12-08 at 21 10 04" src="https://user-images.githubusercontent.com/90475186/206546043-38440ee0-6235-4694-abe1-3298458b7369.png">

* Received file:

<img width="651" alt="Screenshot 2022-12-08 at 21 21 55" src="https://user-images.githubusercontent.com/90475186/206548421-03e2f034-a188-4a60-8577-93a66696e489.png">


* How it looks in Wireshark

    - Sending file bytes:
    
    <img width="540" alt="Screenshot 2022-12-08 at 21 13 20" src="https://user-images.githubusercontent.com/90475186/206546449-d19720eb-8556-4574-b8cf-83bb22b513a5.png">
    
    - Sending file name:
    
    <img width="541" alt="Screenshot 2022-12-08 at 21 08 24" src="https://user-images.githubusercontent.com/90475186/206546542-c092b907-a836-4943-bc61-0d9d8c4ef03d.png">

    - Sending chunked encrypted file:
    
    <img width="539" alt="Screenshot 2022-12-08 at 21 08 39" src="https://user-images.githubusercontent.com/90475186/206546695-1b881283-5d7d-4f72-b036-88b1e4aa98a9.png">

