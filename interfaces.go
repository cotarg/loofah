package main

type Input interface {
	Subscribe() chan string
	Listen() error
}

type Output interface {
	Join(chan string)
	Listen() error
}

type Middleware interface {
	Join(chan string)
	Subscribe() chan string
	Listen() error
}

type StringToJSONMiddleware interface {
	Join(chan string)
	Subscribe() chan JSONLogObject
	Listen() error
}

type JSONMiddleware interface {
	Join(chan JSONLogObject)
	Subscribe() chan JSONLogObject
	Listen() error
}

type JSONToStringMiddleware interface {
	Join(chan JSONLogObject)
	Subscribe() chan string
	Listen() error
}
