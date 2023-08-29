package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Database(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 2.5C14 3.328 11.314 4 8 4s-6-.672-6-1.5S4.686 1 8 1s6 .672 6 1.5z"/><path fill="currentColor" d="M8 5c-3.3 0-6-.7-6-1.5v3C2 7.3 4.7 8 8 8s6-.7 6-1.5v-3C14 4.3 11.3 5 8 5z"/><path fill="currentColor" d="M8 9c-3.3 0-6-.7-6-1.5v3c0 .8 2.7 1.5 6 1.5s6-.7 6-1.5v-3C14 8.3 11.3 9 8 9z"/><path fill="currentColor" d="M8 13c-3.3 0-6-.7-6-1.5v3c0 .8 2.7 1.5 6 1.5s6-.7 6-1.5v-3c0 .8-2.7 1.5-6 1.5z"/>`),
		g.Group(children),
	)
}