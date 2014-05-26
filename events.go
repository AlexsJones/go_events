package main
import (
	"log"
	"time"
	"math/rand"
	"crypto/md5"
	"io"
	"fmt"
)

type State int
const(
	Exiting State = 0
	Running State = 1
)
type EventPusher struct {
	state State
	handler_queue []EventHandler
	fired_queue []string
}
type EventHandlerCallback func(string,uint64)
type EventHandler struct {
	idenitier uint64
	evt_type string
	event_callback EventHandlerCallback
}
func(e* EventPusher) runtime() {
	log.Println("Starting runtime loop...")
	for {
		if len(e.fired_queue) > 0 {
			ev := e.fired_queue[len(e.fired_queue) -1]
			e.fired_queue = e.fired_queue[:len(e.fired_queue) -1]

			for x:=0;x<len(e.handler_queue);x++ {
				handler := e.handler_queue[x]
				if(handler.evt_type == ev) {
					handler.event_callback(ev,handler.idenitier)
				}
			}
		}
		time.Sleep(25 * time.Millisecond)
	}
}
func generateHASH(x string) string {
	h := md5.New()
	io.WriteString(h,x)
	s := fmt.Sprintf("%x",h.Sum(nil))
	return s
}
func generateGUID() uint64 {
	rand.Seed(time.Now().UTC().UnixNano())
	t := uint64(rand.Uint32()) << 32 + uint64(rand.Uint32())
	return t
}
func (e* EventPusher) Init() bool {
	e.state = Running
	go e.runtime()
	return true
}
func (e* EventPusher) GetStatus() State {
	return e.state
}
func (e *EventPusher) AddEventHandler(ev string,func_callback EventHandlerCallback ) uint64{
	id := generateGUID()
	hashed_ev := generateHASH(ev)
	evt := EventHandler { idenitier: id, evt_type:hashed_ev, event_callback: func_callback}
	e.handler_queue = append(e.handler_queue,evt)
	log.Println("Created new event - ",evt)
	return  id
}
func (e *EventPusher) RemoveEventHandler(id uint64) int {
	newslice := make([]EventHandler,len(e.handler_queue) -1)
	for i :=0; i < len(e.handler_queue); i++ {
		current_handler := e.handler_queue[i]
		if(current_handler.idenitier != id) {
			newslice = append(newslice,current_handler)
		}
	}
	e.handler_queue = newslice
	return len(e.handler_queue)
}
func (e *EventPusher) FireEvent(ev string) {
	hashed_ev := generateHASH(ev)
	log.Println("Firing event ",hashed_ev)
	e.fired_queue = append(e.fired_queue,hashed_ev)
}
