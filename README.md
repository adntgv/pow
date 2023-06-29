# pow
proof of work wisdom server
```sh
docker compose build && docker compose up 
```

```
pow-server-1  | 2023/06/27 13:22:53 Server is listening on 0.0.0.0:8000
pow-client-1  | 2023/06/27 13:22:53 Received challenge: 47629
pow-client-1  | 2023/06/27 13:22:53 Found nonce 63007 for challenge 47629
pow-server-1  | 2023/06/27 13:22:53 Verifying nonce 63007 for challenge 47629
pow-client-1  | Received quote: A room without books is like a body without a soul.
```

# Choice of Proof of Work Algorithm
I chose to use a simple proof of work algorithm that is based on the [Hashcash](https://en.wikipedia.org/wiki/Hashcash) algorithm. The algorithm is simple and easy to implement. It is also easy to verify that the client has done the work. The algorithm is also easy to adjust the difficulty by changing the number of leading zeros required.