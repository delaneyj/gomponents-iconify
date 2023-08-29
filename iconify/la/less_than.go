package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LessThan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 6L6 15.219v1.562L26 26v-2.156L9.469 16L26 8.156z"/>`),
		g.Group(children),
	)
}