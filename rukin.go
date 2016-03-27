package main

import (
  "fmt"
  //"github.com/rs/cors"
  "io"
  "net/http"
  "os"
  "path"
  "strings"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {

  // the FormFile function takes in the POST input id file
  file, header, err := r.FormFile("file")

  fileRoot := path.Clean("/www/files/")
  filePath := path.Join(fileRoot, r.URL.Path, header.Filename)
  filePath = path.Clean(filePath);
  if strings.HasPrefix(filePath, fileRoot) {
    err = os.MkdirAll(path.Dir(filePath), os.ModeDir)
    if err != nil {
      fmt.Fprintf(w, "%v", err)
      return
    }
  } else {
    fmt.Fprintf(w, "Path specified is outside file container and cannot be created.")
    return
  }
  fmt.Printf("%s\n", filePath)

  if err != nil {
    fmt.Fprintln(w, "%v", err)
    return
  }

  defer file.Close()

  out, err := os.Create(filePath)
  if err != nil {
    fmt.Fprintf(w, "%v", err)
    return
  }

  defer out.Close()

  // write the content from POST to the file
  _, err = io.Copy(out, file)
  if err != nil {
    fmt.Fprintln(w, err)
  }

  fmt.Fprintf(w, "File uploaded successfully : ")
  fmt.Fprintf(w, filePath)
}

func main() {
  //mux := http.NewServeMux()
  //mux.HandleFunc("/", RequestHandler)
  //handler := cors.Default().Handler(mux)
  //http.ListenAndServe(":8080", handler)

  http.HandleFunc("/", RequestHandler)
  http.ListenAndServe(":8080", nil)
}
