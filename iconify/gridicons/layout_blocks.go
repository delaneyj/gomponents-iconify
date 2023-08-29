package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutBlocks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 7h-2V3a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v2H3a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2v8a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-2h2a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2zm-4 14H7v-8h2a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2H7V3h10v4h-2a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h2v2zm4-4h-6V9h6v8z"/>`),
		g.Group(children),
	)
}