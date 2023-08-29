package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignJustify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4h18v2H3V4Zm0 15h18v2H3v-2Zm0-5h18v2H3v-2Zm0-5h18v2H3V9Z"/>`),
		g.Group(children),
	)
}