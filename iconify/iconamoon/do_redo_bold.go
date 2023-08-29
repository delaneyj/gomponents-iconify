package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoRedoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.884 5.116a1.25 1.25 0 0 0-1.768 1.768l1.768-1.768ZM19 9l.884.884a1.25 1.25 0 0 0 0-1.768L19 9Zm-3.884 2.116a1.25 1.25 0 0 0 1.768 1.768l-1.768-1.768ZM19 18.25a1.25 1.25 0 1 0 0-2.5v2.5ZM15.116 6.884l3 3l1.768-1.768l-3-3l-1.768 1.768Zm3 1.232l-3 3l1.768 1.768l3-3l-1.768-1.768ZM19 7.75H8v2.5h11v-2.5ZM2.75 13c0 2.9 2.35 5.25 5.25 5.25v-2.5A2.75 2.75 0 0 1 5.25 13h-2.5ZM8 7.75A5.25 5.25 0 0 0 2.75 13h2.5A2.75 2.75 0 0 1 8 10.25v-2.5Zm11 8H8v2.5h11v-2.5Z"/>`),
		g.Group(children),
	)
}