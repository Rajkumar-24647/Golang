package main

import (

	"fmt"
	"os"
)

func main() {
//   f, err :=	os.Open("raj.txt")
//   if err != nil {
//? 	log the error
// 	panic(err)
//   }

//  fileInfo, err := f.Stat()

//  if err != nil{
// 	panic(err)
//  }

//  fmt.Println("file name is:", fileInfo.Name())
//  fmt.Println("Is file or folder", fileInfo.IsDir())
//  fmt.Println("The size of the file", fileInfo.Size())
//  fmt.Println("file permission", fileInfo.Mode())
//  fmt.Println("file modified at", fileInfo.ModTime())


//! read file
// f, err := os.Open("raj.txt")
// if err != nil {
// 	panic(err)
// }

// defer f.Close()

// buf := make([]byte, 10)

// d, err := f.Read(buf)
// if err !=nil {
// 	panic(err)
// }

// for i := 0; i < len(buf); i++ {
// 	println("data", d, string(buf[i]))
// }

// fmt.Println("data", d, buf)


//! Another way
//? this is not the efficient way to readFile because it loads the whole file
// data, err := os.ReadFile("raj.txt")
// if err != nil {
// 	panic(err)
// }
// fmt.Println(string(data))


//! Read folders

// dir, err := os.Open("../")
// if err != nil {
// 	panic(err)
// }

// defer dir.Close()
// Info, err := dir.ReadDir(-1)

// for _, fi := range Info {
// 	fmt.Println(fi.Name(), fi.IsDir())
// }


//! create file
// f, err := os.Create("kumar.txt")

// if err != nil {
// 	panic(err)
// }

// defer f.Close()

//? way := 1
// f.WriteString("hii Go")
// f.WriteString("Nice language")
  //? another way
// bytes := []byte("Hello Golang")
// f.Write(bytes)



//! read and write to another file (streaming fashion)

// sourceFile, err := os.Open("raj.txt")
// if err != nil {
// 	panic(err)
// }

// defer sourceFile.Close();

// distFile, err := os.Create("kumar.txt")
// if err != nil {
// 	panic(err)
// }

// defer distFile.Close()

// reader := bufio.NewReader(sourceFile)
// writer :=bufio.NewWriter(distFile)

// for {
// 	b, err :=  reader.ReadByte()
// 	if err!= nil {
// 		if err.Error() != "EOF" {
// 			panic(err)
// 		}

//        break
// 	}

// e := writer.WriteByte(b)

// if e != nil {
// 	panic(e)
// }

// }

// writer.Flush()
// fmt.Println("A new file written successfully")


//! delete  files


err :=  os.Remove("kumar.txt")
if err != nil {
	panic(err)
}

fmt.Println("file deleted successfully")

}