package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paraguay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M30.6 191.3L74.85 34.64C143.1 9.374 223.8 16.21 284.1 36.89l3 143.21l114 11.2l25.5 100.5l54.8-5.2l-27 154.5l-81.8 56.2l-125.2-18l57.7-110.2C192.5 316.9 84.23 263.2 30.6 191.3z"/>`),
		g.Group(children),
	)
}