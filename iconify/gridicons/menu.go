package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Menu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 6v2H3V6h18zM3 18h18v-2H3v2zm0-5h18v-2H3v2z"/>`),
		g.Group(children),
	)
}