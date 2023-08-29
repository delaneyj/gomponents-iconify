package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Previous(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 7a.5.5 0 0 0-.5.5v9a.5.5 0 1 0 1 0v-9a.5.5 0 0 0-.5-.5zm9.354 9.146L12.698 12l4.155-4.146a.5.5 0 0 0-.707-.707l-4.51 4.5a.5.5 0 0 0 0 .707l4.51 4.5a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.708z"/>`),
		g.Group(children),
	)
}