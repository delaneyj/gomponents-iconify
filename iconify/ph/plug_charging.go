package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlugCharging(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 56h-48V16a8 8 0 0 0-16 0v40H96V16a8 8 0 0 0-16 0v40H32.55C26.28 56 24 60.78 24 64a8 8 0 0 0 8 8h16v88a40 40 0 0 0 40 40h32v40a8 8 0 0 0 16 0v-40h32a40 40 0 0 0 40-40V72h16a8 8 0 0 0 0-16Zm-56 128H88a24 24 0 0 1-24-24V72h128v88a24 24 0 0 1-24 24Zm-17.42-60.56a8 8 0 0 1 .91 7.37l-12 32a8 8 0 0 1-15-5.62l8-21.19H112a8 8 0 0 1-7.49-10.81l12-32a8 8 0 1 1 15 5.62l-8 21.19H144a8 8 0 0 1 6.58 3.44Z"/>`),
		g.Group(children),
	)
}