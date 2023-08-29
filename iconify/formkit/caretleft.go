package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Caretleft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m1.38 7.18l4.83-4.15c.7-.6 1.79-.1 1.79.82v8.29c0 .93-1.09 1.42-1.79.82L1.38 8.82c-.5-.43-.5-1.21 0-1.64Z"/>`),
		g.Group(children),
	)
}