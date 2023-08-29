package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Female(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="12" cy="4" r="2" fill="currentColor"/><path fill="currentColor" d="M14.948 7.684A.997.997 0 0 0 14 7h-4a.998.998 0 0 0-.948.684l-2 6l1.775.593L8 18h2v4h4v-4h2l-.827-3.724l1.775-.593l-2-5.999z"/>`),
		g.Group(children),
	)
}