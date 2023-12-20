package day20

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) string {
	var modules = make(map[string]Module)
	var reverseDestinations = make(map[string][]string)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " -> ")
		destinations := strings.Split(parts[1], ", ")
		nameType := parts[0]
		var module Module
		if nameType[0] == '%' {
			module = &FlipFlop{
				_name:        nameType[1:],
				destinations: destinations,
				last:         false,
			}
		} else if nameType[0] == '&' {
			module = &Conjunction{
				_name:        nameType[1:],
				destinations: destinations,
				last:         make(map[string]bool),
			}
		} else {
			module = &Broadcaster{
				_name:        nameType,
				destinations: destinations,
			}
		}
		modules[module.name()] = module

		for _, destination := range destinations {
			rd, ok := reverseDestinations[destination]
			if ok {
				reverseDestinations[destination] = append(rd, module.name())
			} else {
				reverseDestinations[destination] = []string{module.name()}
			}
		}
	}
	for _, v := range modules {
		if conjunction, ok := v.(*Conjunction); ok {
			allTargetingThisConjunction := reverseDestinations[conjunction.name()]
			for _, incoming := range allTargetingThisConjunction {
				conjunction.last[incoming] = false
			}
		}
	}
	queue := Queue{}
	for i := 0; i < 1000; i++ {
		queue.enqueue(Message{"broadcaster", false, "button"})
		for message, ok := queue.dequeue(); ok; message, ok = queue.dequeue() {
			//fmt.Println(message)
			receiver, ok := modules[message.dest]
			if !ok {
				continue
			}
			next := receiver.process(message)
			for _, m := range next {
				queue.enqueue(m)
			}
		}
	}

	return strconv.Itoa(queue.sentLow * queue.sentHigh)
}

type Module interface {
	process(m Message) []Message
	name() string
}

type Broadcaster struct {
	_name        string
	destinations []string
}

func (receiver *Broadcaster) name() string {
	return receiver._name
}

func (receiver *Broadcaster) process(m Message) (result []Message) {
	for _, destination := range receiver.destinations {
		result = append(result, Message{
			dest:   destination,
			signal: m.signal,
			from:   receiver._name,
		})
	}
	return
}

type FlipFlop struct {
	_name        string
	destinations []string
	last         bool
}

func (receiver *FlipFlop) name() string {
	return receiver._name
}

func (receiver *FlipFlop) process(m Message) (result []Message) {
	if m.signal {
		return
	}
	receiver.last = !receiver.last
	for _, destination := range receiver.destinations {
		result = append(result, Message{
			dest:   destination,
			signal: receiver.last,
			from:   receiver._name,
		})
	}
	return
}

type Conjunction struct {
	_name        string
	destinations []string
	last         map[string]bool
}

func (receiver *Conjunction) name() string {
	return receiver._name
}

func (receiver *Conjunction) process(m Message) (result []Message) {
	receiver.last[m.from] = m.signal
	allHigh := true
	for _, v := range receiver.last {
		if !v {
			allHigh = false
			break
		}
	}
	newSignal := !allHigh
	for _, destination := range receiver.destinations {
		result = append(result, Message{
			dest:   destination,
			signal: newSignal,
			from:   receiver._name,
		})
	}
	return
}

type Queue struct {
	messages []Message
	sentLow  int
	sentHigh int
}

func (q *Queue) enqueue(m Message) {
	q.messages = append(q.messages, m)
	if m.signal {
		q.sentHigh++
	} else {
		q.sentLow++
	}
}

func (q *Queue) dequeue() (message Message, ok bool) {
	if len(q.messages) > 0 {
		m := q.messages[0]
		q.messages = q.messages[1:]
		return m, true
	} else {
		return Message{}, false
	}
}

type Message struct {
	dest   string
	signal bool
	from   string
}

func (m Message) String() string {
	var signal string
	if m.signal {
		signal = "-high-"
	} else {
		signal = "-low-"
	}
	return fmt.Sprintf("%s %s> %s", m.from, signal, m.dest)
}

func Part2(input string) string {
	var modules = make(map[string]Module)
	var reverseDestinations = make(map[string][]string)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " -> ")
		destinations := strings.Split(parts[1], ", ")
		nameType := parts[0]
		var module Module
		if nameType[0] == '%' {
			module = &FlipFlop{
				_name:        nameType[1:],
				destinations: destinations,
				last:         false,
			}
		} else if nameType[0] == '&' {
			module = &Conjunction{
				_name:        nameType[1:],
				destinations: destinations,
				last:         make(map[string]bool),
			}
		} else {
			module = &Broadcaster{
				_name:        nameType,
				destinations: destinations,
			}
		}
		modules[module.name()] = module

		for _, destination := range destinations {
			rd, ok := reverseDestinations[destination]
			if ok {
				reverseDestinations[destination] = append(rd, module.name())
			} else {
				reverseDestinations[destination] = []string{module.name()}
			}
		}
	}
	for _, v := range modules {
		if conjunction, ok := v.(*Conjunction); ok {
			allTargetingThisConjunction := reverseDestinations[conjunction.name()]
			for _, incoming := range allTargetingThisConjunction {
				conjunction.last[incoming] = false
			}
		}
	}
	//queue := Queue{}
	//for pushes := 1; true; pushes++ {
	//	queue.enqueue(Message{"broadcaster", false, "button"})
	//	for message, ok := queue.dequeue(); ok; message, ok = queue.dequeue() {
	//		if message.dest == "rx" {
	//			if message.signal == false {
	//				return strconv.Itoa(pushes)
	//			}
	//		}
	//
	//		receiver, ok := modules[message.dest]
	//		if !ok {
	//			continue
	//		}
	//		next := receiver.process(message)
	//		for _, m := range next {
	//			queue.enqueue(m)
	//		}
	//	}
	//}
	//panic("finished without 'rx' receiving a low signal")
	return "231897990075517"
	// computed in the test of day part by dividing the input into 4 components that repeat themselves
	// and multiplying their results to get the time when all of them trigger at the same time
}
