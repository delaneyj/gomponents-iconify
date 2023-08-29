package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 19h4v-2h-4v2zm-4-6h12v-2H6v2zM3 5v2h18V5H3z"/>`),
		g.Group(children),
	)
}