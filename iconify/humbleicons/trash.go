package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 6l.934 13.071A1 1 0 0 0 7.93 20h8.138a1 1 0 0 0 .997-.929L18 6m-6 5v4m8-9H4m4.5 0l.544-1.632A2 2 0 0 1 10.941 3h2.117a2 2 0 0 1 1.898 1.368L15.5 6"/>`),
		g.Group(children),
	)
}