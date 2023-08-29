package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LessThanEqual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 5L6 13.219v1.562L26 23v-2.156L9.469 14L26 7.156zM6 25v2h20v-2z"/>`),
		g.Group(children),
	)
}