This is a simple tutorial of RPC in Go based on the documentation provided in the package [rpc](rpc%5D%28https://pkg.go.dev/net/rpc%29). The example ilustrates the interaction between a server and a client that requests a calc for two numbers based on the operation provided.

### Usage

1. Run `make`
2. In another terminal run `./calc -op <OP> <A> <B>`.

Note:
- OP is a string, and may be one of the followings: "mult", "sum", "sub", "div".
- A and B are two int numbers.

### Example

##### Multiplication

    $ ./calc -op mult 2 3
    2 * 3 = 6 // result

##### Sum

	$  ./calc -op sum 2 3
    2 + 3 = 5 // result
##### Subtraction

	$  ./calc -op sub 2 3
    2 - 3 = -1 // result

##### Division

	$  ./calc -op div 201 3
    201 / 3 = 67 (quociente), 0 (remainder) // result
