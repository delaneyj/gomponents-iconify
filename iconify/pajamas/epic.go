package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Epic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m1.5 8.5l2.5-6h8l2.5 6h-13ZM0 8.569v2.181c0 .966.784 1.75 1.75 1.75H2v.75c0 .966.784 1.75 1.75 1.75h8.5a.75.75 0 0 0 0-1.5h-8.5a.25.25 0 0 1-.25-.25v-.75h10.25a.75.75 0 0 0 0-1.5h-12a.25.25 0 0 1-.25-.25V10h13a1.5 1.5 0 0 0 1.385-2.077l-2.629-6.308A1 1 0 0 0 12.333 1H3.667a1 1 0 0 0-.923.615L.115 7.923A1.498 1.498 0 0 0 0 8.569Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}