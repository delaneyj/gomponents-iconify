package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vuetify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 11.64L7.92 2h8.16L12 11.64m0 10.31L3.27 4.91h4.57L12 14.47l4.16-9.56h4.57L12 21.95Z"/>`),
		g.Group(children),
	)
}