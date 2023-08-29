package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Houzz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.27 0v24h8.05v-7.56h5.36V24h8.05V10.37L6.61 5.75V0H1.27Z"/>`),
		g.Group(children),
	)
}