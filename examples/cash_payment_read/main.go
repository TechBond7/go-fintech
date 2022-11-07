// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"path"

	"github.com/moov-io/fincen"
	"github.com/moov-io/fincen/pkg/batch"
)

func main() {
	buf, err := os.ReadFile(path.Join("data", "samples", "8300_batch.xml"))
	if err != nil {
		return
	}

	r := batch.NewReport(fincen.Form8300)
	err = xml.Unmarshal(buf, r)
	if err != nil {
		return
	}

	buf, err = xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return
	}

	fmt.Println(string(buf))
}
