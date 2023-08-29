package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneRoofing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 16h2v2h-2z" opacity=".3"/><path fill="currentColor" d="M13 18h-2v-2h2v2zm2-4H9v6h6v-6zm4-4.7V4h-3v2.6L12 3L2 12h3l7-6.31L19 12h3l-3-2.7z"/>`),
		g.Group(children),
	)
}