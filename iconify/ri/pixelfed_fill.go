package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PixelfedFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.001 2c5.523 0 10 4.477 10 10s-4.477 10-10 10s-10-4.477-10-10s4.477-10 10-10Zm1.031 6.099h-2.624c-.988 0-1.789.776-1.789 1.733v6.748l2.595-2.471h1.818c1.713 0 3.101-1.345 3.101-3.005s-1.388-3.005-3.1-3.005Z"/>`),
		g.Group(children),
	)
}