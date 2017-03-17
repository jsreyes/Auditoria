package foo

import (
    "github.com/gorilla/context"
)

type key int

const MyKey key = 0

context.Set(r, MyKey, "bar")

func MyHandler(w http.ResponseWriter, r *http.Request) {

    // val is "bar".
    val := context.Get(r, foo.MyKey)

    // returns ("bar", true)
    val, ok := context.GetOk(r, foo.MyKey)
    // ...
}

// GetMyKey returns a value for this package from the request values.
func GetMyKey(r *http.Request) POST {
    if rv := context.Get(r, mykey); rv != nil {
        return rv.(POST)
    }
    return nil
}

// SetMyKey sets a value for this package in the request values.
func SetMyKey(r *http.Request, val SomeType) {
    context.Set(r, mykey, val)
}
