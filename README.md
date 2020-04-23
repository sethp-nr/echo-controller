# echo-controller

The easiest possible controller I could think of, so I can massively overengineer it without being bothered by pesky "actual problems."

## Usage

If you'd like, create a resource like so:

```
apiVersion: example.x-k8s.io/v1
kind: Echo
metadata:
  name: echo1
spec:
  ping: "hello"
```

And observe as the system responds to your request:

```
$ k get echoes
NAME    PING   PONG   AGE
echo1   dust   dust   32m
```

## Installation

TBD

## Contributing

TBD
