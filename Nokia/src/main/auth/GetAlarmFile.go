package main
//
//import (
//   "log"
//   "os"
//   "fmt"
//   "github.com/pkg/sftp"
//   "golang.org/x/crypto/ssh"
//   "path/filepath"
//   "bufio"
//   "strings"
//   "github.com/pkg/errors"
//)
//
//func main() {
//   sshConfig := &ssh.ClientConfig{
//      User:            "toor4nsn",
//      HostKeyCallback: ssh.InsecureIgnoreHostKey(),
//      Auth: []ssh.AuthMethod{
//         ssh.Password("oZPS0POrRieRtu"),
//      },
//   }
//
//   client, err := ssh.Dial("tcp", "192.168.255.129:22", sshConfig)
//   if err != nil {
//      panic("Failed to dial: " + err.Error())
//   }
//   fmt.Println("Successfully connected to ssh server.")
//
//   // open an SFTP session over an existing ssh connection.
//   sftp_client, err := sftp.NewClient(client)
//   if err != nil {
//      log.Fatal(err)
//   }
//   defer sftp_client.Close()
//
//   srcPath := "/tmp/"
//   dstPath := "D:/sftp_temp/"
//   filename := "RawAlarmHistory.txt"
//
//   // Open the source file
//   fmt.Println(srcPath + filename)
//   srcFile, err := sftp_client.Open(srcPath + filename)
//   if err != nil {
//      log.Fatal(err)
//   }
//   defer srcFile.Close()
//
//   // Create the destination file
//   dstFile, err := os.Create(dstPath + filename)
//   if err != nil {
//      log.Fatal(err)
//   }
//   defer dstFile.Close()
//
//   // Copy the file
//   srcFile.WriteTo(dstFile)
//}
