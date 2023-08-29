package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vuejs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m3 6l13 22L29 6H19l-3 5.3L13 6H3zm3.5 2h3.771L16 18l5.729-10H25.5L16 24.1L6.5 8z"/>`),
		g.Group(children),
	)
}