package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19h18v2H3v-2Zm5-6h3l-4 4l-4-4h3V3h2v10Zm10 0h3l-4 4l-4-4h3V3h2v10Z"/>`),
		g.Group(children),
	)
}