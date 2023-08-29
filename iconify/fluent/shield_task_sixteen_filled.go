package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldTaskSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.355 2.147a.5.5 0 0 0-.708 0C6.404 3.388 5.03 4 3.5 4a.5.5 0 0 0-.5.5v3.001c0 3.219 1.641 5.407 4.842 6.473a.499.499 0 0 0 .316 0C11.358 12.908 13 10.72 13 7.501V4.5a.5.5 0 0 0-.5-.5c-1.531 0-2.905-.61-4.145-1.853Zm2.499 4a.5.5 0 0 1 0 .707l-3 3a.5.5 0 0 1-.708 0l-1.5-1.5a.5.5 0 0 1 .708-.708L7.5 8.793l2.646-2.646a.5.5 0 0 1 .708 0Z"/>`),
		g.Group(children),
	)
}