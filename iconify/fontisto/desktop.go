package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Desktop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24.749 0H2.251a2.25 2.25 0 0 0-2.25 2.25v14.999a2.25 2.25 0 0 0 2.25 2.25h8.999l-.75 2.25H7.124a1.125 1.125 0 0 0 0 2.25h12.75a1.125 1.125 0 0 0 0-2.25h-3.375l-.75-2.249h8.999a2.25 2.25 0 0 0 2.25-2.249V2.25A2.25 2.25 0 0 0 24.748 0zm-.751 16.499H2.999V3h20.999z"/>`),
		g.Group(children),
	)
}