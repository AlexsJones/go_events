package main
import(
	"testing"
	"time"
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
	callbackHandler := func(EventTypes,uint64) {
		isCompleted = true
	}
	e.AddEventHandler(Suspending,callbackHandler);
	e.FireEvent(Suspending)
	time.Sleep(1000 * time.Millisecond)
	if(isCompleted != true) {
		t.Error("Callback was not fired within 1000ms")
	}
}
