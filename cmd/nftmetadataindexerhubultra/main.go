// cmd/nftmetadataindexerhubultra/main.go
package main

import (
"flag"
"log"
"os"

"nftmetadataindexerhubultra/internal/nftmetadataindexerhubultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nftmetadataindexerhubultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
