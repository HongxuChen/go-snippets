package main

import (
	"github.com/rickb777/date"
	"go.uber.org/zap"
	"github.com/godbus/dbus"
	"github.com/hackebrot/turtle"
	"fmt"
	"os"
	"mvdan.cc/xurls"
	"path/filepath"
	"github.com/labstack/gommon/log"
)

func dbus_func() {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}
	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	call := obj.Call("org.freedesktop.Notifications.Notify", 0, "", uint32(0),
		"", "Test", "This is a test of the DBus bindings for go.", []string{},
		map[string]dbus.Variant{}, int32(5000))
	if call.Err != nil {
		panic(call.Err)
	}
}

func emoji_func() {
	name := "turtle"
	emoji, ok := turtle.Emojis[name]

	if !ok {
		fmt.Fprintf(os.Stderr, "no emoji found for name: %v\n", name)
		os.Exit(1)
	}

	fmt.Printf("Name: %q\n", emoji.Name)
	fmt.Printf("Char: %s\n", emoji.Char)
	fmt.Printf("Category: %q\n", emoji.Category)
	fmt.Printf("Keywords: %q\n", emoji.Keywords)
}

func xurls_func() {
	var r = xurls.Strict.FindAllString("foo.com is http://foo.com/.", -1)
	fmt.Println(r)
}

func path_func() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}

func main() {
	var sugar = zap.NewExample().Sugar()
	defer sugar.Sync()
	var d_min, d_max = date.Min(), date.Max();
	sugar.Infow("hello world", "d-max", d_max, "d-min", d_min)
	dbus_func()
	emoji_func()
	xurls_func()
	path_func()
}
