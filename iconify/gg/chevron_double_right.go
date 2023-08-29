package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5.636 7.757L7.05 6.343L12.707 12L7.05 17.657l-1.414-1.414L9.88 12L5.636 7.757Z"/><path d="m12.707 6.343l-1.414 1.414L15.536 12l-4.243 4.243l1.414 1.414L18.364 12l-5.657-5.657Z"/></g>`),
		g.Group(children),
	)
}