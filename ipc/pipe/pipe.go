package pipe

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	rand2 "math/rand"
	"time"
)

func Read(reader *io.PipeReader)  {
	//buffer := bytes.NewBuffer(make([]byte,0))
	bytes := make([]byte,8)
	for {

		n, err := reader.Read(bytes)
		if err != nil {
			log.Fatal(err)
		}

		if n > 0 {
			fmt.Printf("%d\n",binary.BigEndian.Uint64(bytes))
		}
	}
}

func Write(writer *io.PipeWriter)  {
	for {
		rand := rand2.New(rand2.NewSource(time.Now().Unix()))
		n := rand.Uint64()

		bytes := make([]byte,8)
		binary.BigEndian.PutUint64(bytes,n)

		if _,err := writer.Write(bytes); err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Second)
	}
}