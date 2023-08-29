package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tshirt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1009 265l-72 72q-15 15-36.5 15T864 337l-32-31v527q0 26-18.5 45T768 897H256q-27 0-45.5-19T192 833V306l-32 31q-15 15-36.5 15T87 337l-72-72Q0 250 0 229t15-36L192 16q15-16 64-16h128q0 37 33.5 50.5T512 64t94.5-13.5T640 0h128q49 0 64 16l177 177q15 15 15 36t-15 36z"/>`),
		g.Group(children),
	)
}