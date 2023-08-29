package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21h15v1c0 .54-.44 1-1 1H2c-.55 0-1-.46-1-1v-1m19.5 2.31L16 18.83V19H1v-2h13.17l-2-2H1c0-3.24 2.46-5.17 5.38-5.79L.68 3.5L2.1 2.1L13 13l2 2l6.9 6.9l-1.4 1.41M10.17 13l-2-2c-1.42.06-3.52.56-4.55 2h6.55M23 5h-5V1h-2v4h-5l.23 2h9.56l-1 9.97l1.83 1.83L23 5Z"/>`),
		g.Group(children),
	)
}