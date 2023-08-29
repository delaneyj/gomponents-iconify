package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoUndoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.884 6.884a1.25 1.25 0 1 0-1.768-1.768l1.768 1.768ZM5 9l-.884-.884a1.25 1.25 0 0 0 0 1.768L5 9Zm2.116 3.884a1.25 1.25 0 1 0 1.768-1.768l-1.768 1.768ZM5 15.75a1.25 1.25 0 0 0 0 2.5v-2.5ZM7.116 5.116l-3 3l1.768 1.768l3-3l-1.768-1.768Zm-3 4.768l3 3l1.768-1.768l-3-3l-1.768 1.768ZM5 10.25h11v-2.5H5v2.5ZM18.75 13A2.75 2.75 0 0 1 16 15.75v2.5c2.9 0 5.25-2.35 5.25-5.25h-2.5ZM16 10.25A2.75 2.75 0 0 1 18.75 13h2.5c0-2.9-2.35-5.25-5.25-5.25v2.5Zm-11 8h11v-2.5H5v2.5Z"/>`),
		g.Group(children),
	)
}