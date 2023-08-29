package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 11h8a1 1 0 0 0 .77-.37A1 1 0 0 0 17 9.8l-1-5a1 1 0 0 0-1-.8H9a1 1 0 0 0-1 .8l-1 5a1 1 0 0 0 .21.83A1 1 0 0 0 8 11Zm1.82-5h4.36l.6 3H9.22ZM22 13.8a1 1 0 0 0-1-.8h-6a1 1 0 0 0-1 .8l-1 5a1 1 0 0 0 .21.83A1 1 0 0 0 14 20h8a1 1 0 0 0 .77-.37a1 1 0 0 0 .23-.83ZM15.22 18l.6-3h4.36l.6 3ZM9 13H3a1 1 0 0 0-1 .8l-1 5a1 1 0 0 0 .21.83A1 1 0 0 0 2 20h8a1 1 0 0 0 .77-.37a1 1 0 0 0 .23-.83l-1-5a1 1 0 0 0-1-.8Zm-5.78 5l.6-3h4.36l.6 3Z"/>`),
		g.Group(children),
	)
}