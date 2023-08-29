package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.25 10a1.25 1.25 0 1 0-2.5 0h2.5Zm-14 0a1.25 1.25 0 1 0-2.5 0h2.5Zm13.866 2.884a1.25 1.25 0 0 0 1.768-1.768l-1.768 1.768ZM12 3l.884-.884a1.25 1.25 0 0 0-1.768 0L12 3Zm-9.884 8.116a1.25 1.25 0 0 0 1.768 1.768l-1.768-1.768ZM7 22.25h10v-2.5H7v2.5ZM20.25 19v-9h-2.5v9h2.5Zm-14 0v-9h-2.5v9h2.5Zm15.634-7.884l-9-9l-1.768 1.768l9 9l1.768-1.768Zm-10.768-9l-9 9l1.768 1.768l9-9l-1.768-1.768ZM17 22.25A3.25 3.25 0 0 0 20.25 19h-2.5a.75.75 0 0 1-.75.75v2.5Zm-10-2.5a.75.75 0 0 1-.75-.75h-2.5A3.25 3.25 0 0 0 7 22.25v-2.5Z"/>`),
		g.Group(children),
	)
}