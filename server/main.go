package main

import (
    // "fmt"
    "log"
    "net"
    "net/http"
    "net/rpc"
)

type API int

type Item struct {
  Title   string
  Body    string
}

var database []Item

func (a *API) GetDB(empty string, reply *[]Item) error {
  *reply = database
  return nil
}

func (a *API) GetByName(title string, reply *Item) error {
  var getItem Item
  for _, val := range database {
    if val.Title == title {
      getItem = val
    }
  }
  *reply = getItem
  return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
  database = append(database, item);
  *reply = item;
  return nil
}

func (a *API) EditItem(editItem Item, reply *Item) error {
  var changedItem Item;

  for idx, val := range database {
    if val.Title == editItem.Title {
      database[idx] = Item{editItem.Title, editItem.Body};
      changedItem = database[idx];
    }
  }

  *reply = changedItem;
  return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error  {
  var deleteItem Item
  for idx, val := range database {
    if val.Title == item.Title {
      database = append(database[:idx], database[idx+1:]...)
      deleteItem = item;
      break;
    }
  }
  *reply = deleteItem;
  return nil;
}

func main() {

  api := new(API)
  err := rpc.Register(api)
  if err != nil {
    log.Fatal("error registering API", err)
  }

  rpc.HandleHTTP()

  listener, err := net.Listen("tcp", ":4040")

  if err != nil {
    log.Fatal("Listener error", err)
  }
  log.Printf("serving rpc on port %d", 4040)
  http.Serve(listener, nil)

  if err != nil {
    log.Fatal("error serving: ", err)
  }


  // fmt.Println("initial database: ", database)
  // a := Item{"first", "a test item"}
  // b := Item{"second", "a second item"}
  // c := Item{"third", "a third item"}

  // AddItem(a)
  // AddItem(b)
  // AddItem(c)
  // fmt.Println("second database: ", database)

  // DeleteItem(b)
  // fmt.Println("third database: ", database)

  // EditItem("third", Item{"fourth", "a new item"})
  // fmt.Println("fourth database: ", database)

  // x := GetByName("fourth")
  // y := GetByName("first")
  // fmt.Println(x, y)

  // handleRequests()
}