package main

import "github.com/Azure/azure-sdk-for-go/storage"
import "fmt"
import "time"

var (
	tableCli storage.TableServiceClient
)

const (
	account = "account"
	key = "key"
	fullmetadata = "application/json;odata=fullmetadata"
	tablename = "table name"
)

func main(){
	//insert("1", "3")
	//query()
}

func insert(partitionkey string, rowkey string){

	client, err := storage.NewBasicClient(account,key)

	if err != nil {
		fmt.Printf("%s: \n", err)
	}

	tableCli = client.GetTableService()

	fmt.Println(tableCli)

	table := tableCli.GetTableReference(tablename)

	entity := table.GetEntityReference(partitionkey, rowkey)

	props := map[string]interface{}{
		"AmountDue":      200.23,
		"CustomerCode":   "123",
		"CustomerSince":  time.Date(1992, time.December, 20, 21, 55, 0, 0, time.UTC),
		"IsActive":       true,
		"NumberOfOrders": int64(255),
	}

	entity.Properties = props

	err = entity.Insert(fullmetadata, nil)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Inserted! ")
	}
}

func query(){
	s
	client, err := storage.NewBasicClient(account,key)

	if err != nil {
		fmt.Printf("%s: \n", err)
	}

	tableCli = client.GetTableService()

	fmt.Println(tableCli)

	table := tableCli.GetTableReference(tablename)

	// timeout, metatadalevel, options
	entities, err := table.QueryEntities(30, fullmetadata, nil)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(entities.Entities[0])
	}
}
