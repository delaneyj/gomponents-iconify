package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.5 5h-4a.5.5 0 0 0 0 1H21v12H10.707l1.646-1.646a.5.5 0 0 0-.707-.707l-2.5 2.5a.5.5 0 0 0 0 .707l2.5 2.5a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.708L10.708 19h10.794a.5.5 0 0 0 .499-.5v-13a.5.5 0 0 0-.5-.5zm-16 13H3V6h10.293l-1.647 1.646a.5.5 0 1 0 .707.708l2.5-2.5a.5.5 0 0 0 0-.707l-2.5-2.5a.5.5 0 0 0-.707.707L13.293 5H2.499A.5.5 0 0 0 2 5.5v13a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}