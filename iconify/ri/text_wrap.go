package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextWrap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 18h1.5a2.5 2.5 0 0 0 0-5H3v-2h13.5a4.5 4.5 0 1 1 0 9H15v2l-4-3l4-3v2ZM3 4h18v2H3V4Zm6 14v2H3v-2h6Z"/>`),
		g.Group(children),
	)
}