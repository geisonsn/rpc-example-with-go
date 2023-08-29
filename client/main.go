package main

import (
	"errors"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	var operation string

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "op",
				Value:       "sum",
				Usage:       "calculating two numbers",
				Destination: &operation,
			},
		},
		Action: func(cCtx *cli.Context) error {
			if operation == "" {
				log.Fatal("No operation provided")
			}

			a, b, err := getParams(cCtx)
			if err != nil {
				log.Fatalf("Fail to get the params a and b: %s", err.Error())
			}

			if err := CalcFromServer(a, b, operation); err != nil {
				log.Fatal(err)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func CalcFromServer(a, b int, operation string) error {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		return err
	}

	if operation == "sum" {
		return sum(a, b, client)
	} else if operation == "sub" {
		return sub(a, b, client)
	} else if operation == "mult" {
		return mult(a, b, client)
	} else if operation == "div" {
		return div(a, b, client)
	} else {
		return errors.New("invalid operation")
	}
}

func getParams(cCtx *cli.Context) (int, int, error) {
	arga := cCtx.Args().Get(0)
	argb := cCtx.Args().Get(1)

	a, err := strconv.Atoi(arga)
	if err != nil {
		return 0, 0, errors.New("fail to convert a to an int")
	}

	b, err := strconv.Atoi(argb)
	if err != nil {
		return 0, 0, errors.New("fail to convert b to an int")
	}
	return a, b, nil
}

func sum(a, b int, client *rpc.Client) error {
	// Synchronous call
	args := &Args{a, b}
	var r int
	err := client.Call("Arith.Sum", args, &r)
	if err != nil {
		return err
	}

	fmt.Printf("%d + %d = %d\n", args.A, args.B, r)

	return nil
}

func sub(a, b int, client *rpc.Client) error {
	// Synchronous call
	args := &Args{a, b}
	var r int
	err := client.Call("Arith.Subtract", args, &r)
	if err != nil {
		return err
	}

	fmt.Printf("%d - %d = %d\n", args.A, args.B, r)

	return nil
}

func mult(a, b int, client *rpc.Client) error {
	// Synchronous call
	args := &Args{a, b}
	var r int
	err := client.Call("Arith.Multiply", args, &r)
	if err != nil {
		return err
	}

	fmt.Printf("%d * %d = %d\n", args.A, args.B, r)

	return nil
}

func div(a, b int, client *rpc.Client) error {
	// Synchronous call
	args := &Args{a, b}
	var q Quotient
	err := client.Call("Arith.Divide", args, &q)
	if err != nil {
		return err
	}

	fmt.Printf("%d / %d = %d (quociente), %d (remainder)\n", args.A, args.B, q.Quo, q.Rem)
	return nil
}
