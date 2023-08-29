package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FruitWatermelon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.4 16.4c3.4-3.4 3.4-8.9 0-12.2L4.2 16.4c3.3 3.4 8.8 3.4 12.2 0M16 7c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1m0 4c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1m-4 0c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1m0 4c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1m-4 2c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1m10.6 1.6c-4.6 4.6-12 4.6-16.6 0l1.4-1.4c3.8 3.8 9.9 3.8 13.7 0c3.8-3.8 3.8-9.9 0-13.7L18.6 2c4.5 4.6 4.5 12 0 16.6Z"/>`),
		g.Group(children),
	)
}