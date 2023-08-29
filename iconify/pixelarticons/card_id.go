package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardId(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4h20v16H2V4zm2 2v4h16V6H4zm16 6H10v2h10v-2zm0 4h-4v2h4v-2zm-6 2v-2H4v2h10zM4 14h4v-2H4v2z"/>`),
		g.Group(children),
	)
}