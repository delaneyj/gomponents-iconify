package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hexlet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.775 3.042L12 0L4.225 3.042L12 5.746ZM16.732 7.1v6.422H7.268V7.1L4.563 6.085V24h2.705v-7.775h9.464V24h2.705V6.085Z"/>`),
		g.Group(children),
	)
}