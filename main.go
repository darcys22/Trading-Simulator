package main

import (
	//"context"
	//"log"
	//"time"
	"fmt"

	"math"
	"math/rand"
	//pb "github.com/darcys22/godbledger/proto"
	//"google.golang.org/grpc"
)

const (
	address       = "localhost:50051"
	iterationDays = 365
	sdBPS         = 500
	decimalsBPS   = 100
	tradProb      = 10
	startPrice    = 100
)

// Account holds the name and balance
type Trade struct {
	amount int
	price  int
}

func main() {

	rand.Seed(42)

	//positions = []Trade{}
	price := startPrice

	for day := 1; day <= iterationDays; day++ {
		price = price * (1 + math.Round(rand.NormFloat64()*sdBPS)/decimalsBPS)
		fmt.Println("Price: " + price)
		if rand.Intn(10) == 1 {
			if rand.Intn(2) == 1 {
				fmt.Println("Buy")
			} else {
				fmt.Println("Sell")
			}
		}
	}

	// Set up a connection to the server.
	//conn, err := grpc.Dial(address, grpc.WithInsecure())
	//if err != nil {
	//log.Fatalf("did not connect: %v", err)
	//}
	//defer conn.Close()
	//client := pb.NewTransactorClient(conn)

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()

	//date := "2011-03-15"
	//desc := "Whole Food Market"

	//transactionLines := make([]*pb.LineItem, 2)

	//line1Account := "Expenses:Groceries"
	//line1Desc := "Groceries"
	//line1Amount := int64(7500)

	//transactionLines[0] = &pb.LineItem{
	//Accountname: line1Account,
	//Description: line1Desc,
	//Amount:      line1Amount,
	//}

	//line2Account := "Assets:Checking"
	//line2Desc := "Groceries"
	//line2Amount := int64(-7500)

	//transactionLines[1] = &pb.LineItem{
	//Accountname: line2Account,
	//Description: line2Desc,
	//Amount:      line2Amount,
	//}

	//req := &pb.TransactionRequest{
	//Date:        date,
	//Description: desc,
	//Lines:       transactionLines,
	//}
	//r, err := client.AddTransaction(ctx, req)
	//if err != nil {
	//log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Version: %s", r.GetMessage())
}
