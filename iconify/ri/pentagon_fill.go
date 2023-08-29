package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PentagonFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 .7l10.747 7.808l-4.105 12.634H5.358L1.253 8.508L12 .7Z"/>`),
		g.Group(children),
	)
}