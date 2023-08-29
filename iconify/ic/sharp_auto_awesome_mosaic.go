package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpAutoAwesomeMosaic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21h8V3H3v18zM21 3h-8v8h8V3zm-8 18h8v-8h-8v8z"/>`),
		g.Group(children),
	)
}