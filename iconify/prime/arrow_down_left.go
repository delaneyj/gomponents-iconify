package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.54 17.7a.75.75 0 0 0 0-1.5H8.86l8.62-8.62a.75.75 0 1 0-1.06-1.06L7.8 15.14V8.46a.75.75 0 0 0-1.5 0V17a.75.75 0 0 0 .06.29a.76.76 0 0 0 .69.46Z"/>`),
		g.Group(children),
	)
}