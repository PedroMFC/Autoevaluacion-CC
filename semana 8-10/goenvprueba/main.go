package main

import (
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
	"golang.org/x/net/context"
	"time"
    "github.com/joho/godotenv"
    "log"
	"os"
	// Manualmente desde la terminal
	// $> go get go.etcd.io/etcd/clientv3@release-3.4
	"go.etcd.io/etcd/clientv3"
)

func main() {
	/*
	 * Probamos .env
	 */ 
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")
  
	// now do something with s3 or whatever
	log.Print("Prueba .env:" + s3Bucket + " " + secretKey)
  
	log.Print("-----------")

	/*
	 * Probamos etcd
	 */ 

	// Creamos el cliente
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})

	// Hacemos una petición de put
	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	resp, err := cli.Put(ctx, "sample_key", "sample_value")
	if err != nil {
		switch err {
		case context.Canceled:
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}

	log.Print(resp)

	// Hacemos una petición de get
	resp1, err1 := cli.Get(ctx, "keyPrueba")
	if err1 != nil {
		// Hacer algo ....
	}

	log.Print(string(resp1.Kvs[0].Value))

	for _, ev := range resp1.Kvs {
		log.Printf("Otra opción -> %s : %s\n", ev.Key, ev.Value)
	}

	resp1, err1 = cli.Get(ctx, "sample_key")
	if err1 != nil {
		// Hacer algo ....
	}

	log.Print(string(resp1.Kvs[0].Value))

	for _, ev := range resp1.Kvs {
		log.Printf("Otra opción -> %s : %s\n", ev.Key, ev.Value)
	}
	
}