package et

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M17 6.5a.5.5 0 0 0-1 0V15H6.5a.5.5 0 0 0 0 1H16v9.469a.5.5 0 0 0 1 0V16h9.469a.5.5 0 0 0 0-1H17V6.5z"/><path d="M31.562 15h-2.613C28.461 8.63 23.37 3.539 17 3.051V.5a.5.5 0 0 0-1 0V3C9.17 3 3.565 8.299 3.051 15H.562a.5.5 0 0 0 0 1H3c0 7.168 5.832 13 13 13v2.5a.5.5 0 0 0 1 0v-2.551C23.701 28.435 29 22.83 29 16h2.562a.5.5 0 0 0 0-1zM16 28C9.383 28 4 22.617 4 16S9.383 4 16 4s12 5.383 12 12s-5.383 12-12 12z"/></g>`),
		g.Group(children),
	)
}