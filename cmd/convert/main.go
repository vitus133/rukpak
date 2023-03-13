package main

import (
	"fmt"
	"os"

	convert "github.com/operator-framework/rukpak/internal/convert"
)

func main() {
	fsys := os.DirFS("/home/vgrinber/workspc/no-olm/bundle")
	objects, err := convert.RegistryV1ToPlain(fsys)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print(objects.String())
}
