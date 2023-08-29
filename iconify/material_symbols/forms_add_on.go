package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormsAddOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 20.975v-3h-3v-2h3v-3h2v3h3v2h-3v3h-2ZM3 18v-2h2v2H3Zm4 0v-2h4.075q-.075.525-.063 1t.088 1H7Zm-4-4v-2h2v2H3Zm4 0v-2h6.65q-.575.4-1.038.9T11.8 14H7Zm-4-4V8h2v2H3Zm4 0V8h12v2H7ZM3 6V4h2v2H3Zm4 0V4h12v2H7Z"/>`),
		g.Group(children),
	)
}