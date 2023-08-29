package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 6H8V2H6v4H2v2h4v4h2V8h4V6zm18 4V4h-6v2h-8v2h8v2h2v14h-2v2H10v-2H8v-8H6v8H4v6h6v-2h14v2h6v-6h-2V10zM8 28H6v-2h2zm20 0h-2v-2h2zM26 6h2v2h-2z"/>`),
		g.Group(children),
	)
}