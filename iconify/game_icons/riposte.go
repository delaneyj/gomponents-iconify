package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Riposte(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M264.8 14.86L243.5 130.7L133.8 98.86l64.3 83.04C8.624 261.7 55.82 556 383 486.7C274.4 499.9 66.42 382 240.6 284.1l7.2 73.1L305 258l114.4 26.8l-74.7-85.2l80.3-69.2l-111 5.5z"/>`),
		g.Group(children),
	)
}