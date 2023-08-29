package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckmarkFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M40.317 12.695a1.5 1.5 0 0 1-.012 2.122l-22.25 22a1.5 1.5 0 0 1-2.101.008l-9.25-9a1.5 1.5 0 1 1 2.092-2.15l8.195 7.974l21.204-20.966a1.5 1.5 0 0 1 2.122.012Z"/>`),
		g.Group(children),
	)
}