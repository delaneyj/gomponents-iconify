package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendDownFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M22 17.01V13a1 1 0 0 0-1.707-.707L19 13.586l-5.293-5.293a1 1 0 0 0-1.414 0L9 11.586L3.707 6.293a1 1 0 0 0-1.414 1.414l6 6a1 1 0 0 0 1.414 0L13 10.414L17.586 15l-1.293 1.293A1 1 0 0 0 17 18h4l.048-.001a.997.997 0 0 0 .952-.99Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}