# cccurl
A minimal curl-like HTTP client written in Go.

Built as part of the
[Build Your Own curl](https://codingchallenges.fyi/challenges/challenge-curl/#step-1)
 challenge to learn HTTP and Go fundamentals.

Current Features:
- Parses a single URL from the command line
- Validates http scheme
- Extracts host, port (defaults to 80), and path (including query)
- Prints the HTTP/1.1 request that would be sent

### Usage
`go run . http://example.com`

Example output:

```
connecting to example.com
Sending request GET / HTTP/1.1
Host: example.com
Accept: */*
```

### Goals
The focus is on understanding:

- How HTTP/1.1 request lines are constructed
- How URLs are structured and parsed
- Learn idiomatic Go

### Next Steps
Step 2: 
- Open a TCP connection with net.Dial
- Send the constructed request
- Read and print the response
- Add flags (-X, -H, -d, -v)
