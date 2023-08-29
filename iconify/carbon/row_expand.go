package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RowExpand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 18h11v6.172l-2.586-2.586L11 23l5 5l5-5l-1.414-1.414L17 24.172V18h11v-2H4v2zM26 4H6a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2zm0 6H6V6h20z"/>`),
		g.Group(children),
	)
}