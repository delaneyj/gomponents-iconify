package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clouds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M16 7.5a2.5 2.5 0 0 1-1.456 2.272a3.513 3.513 0 0 0-.65-.824a1.5 1.5 0 0 0-.789-2.896a.5.5 0 0 1-.627-.421a3 3 0 0 0-5.22-1.625a5.587 5.587 0 0 0-1.276.088a4.002 4.002 0 0 1 7.392.91A2.5 2.5 0 0 1 16 7.5z"/><path d="M7 5a4.5 4.5 0 0 1 4.473 4h.027a2.5 2.5 0 0 1 0 5H3a3 3 0 0 1-.247-5.99A4.502 4.502 0 0 1 7 5zm3.5 4.5a3.5 3.5 0 0 0-6.89-.873a.5.5 0 0 1-.51.375A2 2 0 1 0 3 13h8.5a1.5 1.5 0 1 0-.376-2.953a.5.5 0 0 1-.624-.492V9.5z"/></g>`),
		g.Group(children),
	)
}