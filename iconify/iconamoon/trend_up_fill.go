package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendUpFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M22 6.99V11a1 1 0 0 1-1.707.707L19 10.414l-5.293 5.293a1 1 0 0 1-1.414 0L9 12.414l-5.293 5.293a1 1 0 0 1-1.414-1.414l6-6a1 1 0 0 1 1.414 0L13 13.586L17.586 9l-1.293-1.293A1 1 0 0 1 17 6h4l.048.001a.996.996 0 0 1 .952.99Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}