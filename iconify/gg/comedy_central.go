package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComedyCentral(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M10.544 19a7 7 0 1 0-4.95-11.95L3.475 4.93l-.019.018A9.969 9.969 0 0 1 10.545 2c5.522 0 10 4.477 10 10s-4.478 10-10 10a9.969 9.969 0 0 1-7.072-2.929l2.122-2.121a6.978 6.978 0 0 0 4.95 2.05Z"/><path d="M10.544 14c.594 0 1.126-.258 1.493-.668l2.122 2.122a5 5 0 1 1 0-6.909l-2.122 2.123A2 2 0 1 0 10.545 14Z"/></g>`),
		g.Group(children),
	)
}