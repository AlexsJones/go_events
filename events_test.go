package main
import(
	"testing"
	"time"
	"fmt"
)

func TestInitialisation(t *testing.T) {
	e := EventPusher{}
	e.Init()
	if(e.GetStatus() != Running) {
		t.Error("Could not initialize event pusxher")
	}
}
func TestSubscription(t *testing.T) {
	e := EventPusher{}
	e.Init()
	isCompleted := false
	callbackHandler := func(string,uint64) {
		isCompleted = true
	}
	id := e.AddEventHandler("FOO",callbackHandler)
	fmt.Println("Added event handler with ID ",id)
	e.FireEvent("FOO")
	time.Sleep(1000 * time.Millisecond)
	if(isCompleted != true) {
		t.Error("Callback was not fired within 1000ms")
	}
}
func TestUnsubscription(t *testing.T) {
	e := EventPusher{}
	e.Init()
	isCompleted := false
	callbackHandler := func(string,uint64) { isCompleted = true }
	id := e.AddEventHandler("BAR",callbackHandler)
	num := e.RemoveEventHandler(id)
	if(num > 0) {
		t.Error("Something went wrong. The remaining number of handlers is greater than 0!")
	}
}
