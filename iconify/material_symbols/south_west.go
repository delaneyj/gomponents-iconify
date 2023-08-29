package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SouthWest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19V9h2v6.6L18.6 4L20 5.4L8.4 17H15v2H5Z"/>`),
		g.Group(children),
	)
}