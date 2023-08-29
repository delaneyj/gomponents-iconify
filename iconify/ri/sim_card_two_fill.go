package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimCardTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 2h10l4.707 4.707a1 1 0 0 1 .293.707V21a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1Zm8 16v-8H8v2h3v6h2Zm-5-5v2h2v-2H8Zm6 0v2h2v-2h-2Zm0-3v2h2v-2h-2Zm-6 6v2h2v-2H8Zm6 0v2h2v-2h-2Z"/>`),
		g.Group(children),
	)
}