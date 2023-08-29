package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberZeroDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 40v176a16 16 0 0 1-16 16H56a16 16 0 0 1-16-16V40a16 16 0 0 1 16-16h144a16 16 0 0 1 16 16Z" opacity=".2"/><path d="M183.25 63.2C170.25 42.79 151.15 32 128 32S85.75 42.79 72.75 63.2C62 80.18 56 103.19 56 128s6 47.82 16.75 64.8c13 20.41 32.1 31.2 55.25 31.2s42.25-10.79 55.25-31.2c10.8-17 16.75-40 16.75-64.8s-5.95-47.82-16.75-64.8ZM128 208c-38.68 0-56-40.18-56-80s17.32-80 56-80s56 40.18 56 80s-17.32 80-56 80Z"/></g>`),
		g.Group(children),
	)
}