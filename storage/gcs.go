// Sample storage-quickstart creates a Google Cloud Storage bucket.
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
	// "google.golang.org/api/iterator"

)

func main() {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	// projectID := "personal-exploration-363811"
	_ = "personal-exploration-363811"

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Sets the name for the new bucket.
	bucketName := "yash-test-bucket-1"

	// Creates a Bucket instance.
	bucket := client.Bucket(bucketName)

	// Creates the new bucket.
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	// if err := bucket.Create(ctx, projectID, nil); err != nil {
	// 	log.Fatalf("Failed to create bucket: %v", err)
	// }

	// fmt.Printf("Bucket %v created.\n", bucketName)
	attrs, err := bucket.Attrs(ctx)
	if err != nil {
		log.Fatal("unable to list attributes")
	}
	fmt.Printf("bucket %s\n, created at %s\n, is located in %s\n with storage class %s\n",
		attrs.Name, attrs.Created, attrs.Location, attrs.StorageClass)

	obj := bucket.Object("data/pics")
	// Write something to obj.
	// w implements io.Writer.
	// w := obj.NewWriter(ctx)
	// // Write some text to obj. This will either create the object or overwrite whatever is there already.
	// if _, err := fmt.Fprintf(w, "This object contains text.\n"); err != nil {
	// 	// TODO: Handle error.
	// }
	// // Close, just like writing a file.
	// if err := w.Close(); err != nil {
	// 	// TODO: Handle error.
	// }

	// Read it back.
	r, err := obj.NewReader(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer r.Close()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		// TODO: Handle error.
	}
	// Prints "This object contains text."
	objAttrs, err := obj.Attrs(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	fmt.Printf("object %s \n has size %d \n and can be read using %s \n",
		objAttrs.Name, objAttrs.Size, objAttrs.MediaLink)
}
