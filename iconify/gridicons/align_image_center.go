package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignImageCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5h18v2H3V5zm0 14h18v-2H3v2zm5-4h8V9H8v6z"/>`),
		g.Group(children),
	)
}