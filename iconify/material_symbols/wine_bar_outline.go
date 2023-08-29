package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WineBarOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21v-2h3v-4.1q-2.15-.35-3.575-2T6 9V3h12v6q0 2.25-1.425 3.9T13 14.9V19h3v2H8Zm4-8q1.4 0 2.45-.85t1.4-2.15h-7.7q.35 1.3 1.4 2.15T12 13ZM8 8h8V5H8v3Zm4 5Z"/>`),
		g.Group(children),
	)
}