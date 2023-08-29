package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditorInsertmore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 7V3H3v4h14zM6 11V9H3v2h3zm6 0V9H8v2h4zm5 0V9h-3v2h3zm0 6v-4H3v4h14z"/>`),
		g.Group(children),
	)
}