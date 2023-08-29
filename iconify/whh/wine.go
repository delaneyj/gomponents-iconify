package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M384 632v264h128q26 0 45 18.5t19 45t-19 45.5t-45 19H128q-27 0-45.5-19T64 959.5t18.5-45T128 896h128V632Q145 605 72.5 499.5T0 256Q0 110 82 0h476q82 110 82 256q0 138-73 243.5T384 632zM524 64H116q-52 86-52 192q0 67 21 128h469q22-61 22-128q0-106-52-192z"/>`),
		g.Group(children),
	)
}