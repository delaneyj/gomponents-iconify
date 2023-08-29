package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpFilterBAndW(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 3H3v18h18V3zm-2 16l-7-8v8H5l7-8V5h7v14z"/>`),
		g.Group(children),
	)
}