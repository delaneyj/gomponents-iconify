package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReligiousJewishEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M11 8H7.1l-1.6 3l-1.6-3H0l1.95-2.5L0 3h3.9l1.6-3l1.6 3H11L9.05 5.5L11 8z" fill="currentColor"/>`),
		g.Group(children),
	)
}