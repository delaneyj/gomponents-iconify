package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvOptionsEditChannels(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 16v-2h2v2H2Zm0-4v-2h2v2H2Zm0-4V6h2v2H2Zm4 8v-2h4v2H6Zm0-4v-2h13v2H6Zm0-4V6h13v2H6Zm9.95 12l-4.25-4.25l1.425-1.425l2.825 2.825l5.65-5.65l1.4 1.45L15.95 20Z"/>`),
		g.Group(children),
	)
}