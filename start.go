package main
import (
	"time"
	"log"
	"math/rand"
)
func callbackHandler(evt EventTypes,id uint64) {
	log.Println("Callback received from fired event  ",id, evt)
}

func main() {

	e := EventPusher{}
	e.Init()
	rand.Seed(time.Now().UTC().UnixNano())
	d := func( x int32) time.Duration{
		t := time.Duration(rand.Int31n(x))
		return  time.Duration(t * time.Second)
	}

	e.AddEventHandler(Suspending,callbackHandler)
	e.AddEventHandler(WakeUp,callbackHandler)
	e.AddEventHandler(Sleep,callbackHandler)

	for {
		log.Println("waiting...")
		delay := d(5)
		log.Println("With delay- ",delay)
		time.Sleep(delay)

		//add some arbitrary events
		i := EventTypes(rand.Intn(2))
		e.FireEvent(i)
	}
}
