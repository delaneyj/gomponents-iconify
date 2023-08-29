package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Constraint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 15v-1.5a2.5 2.5 0 0 0-5 0V15a1.075 1.075 0 0 0-1 1v4a1.075 1.075 0 0 0 1 1h5a1.075 1.075 0 0 0 1-1v-4a1.075 1.075 0 0 0-1-1Zm-1 0h-3v-1.5a1.5 1.5 0 0 1 3 0Zm2-10H2V3h20Zm0 4H2V7h20Zm-9 4H2v-2h11Zm0 4H2v-2h11Zm0 4H2v-2h11Z"/>`),
		g.Group(children),
	)
}