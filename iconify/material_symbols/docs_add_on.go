package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocsAddOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 20.975v-3h-3v-2h3v-3h2v3h3v2h-3v3h-2ZM4 18v-2h7.075q-.075.525-.063 1t.088 1H4Zm0-4v-2h9.65q-.575.4-1.038.9T11.8 14H4Zm0-4V8h15v2H4Zm0-4V4h15v2H4Z"/>`),
		g.Group(children),
	)
}