package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkStrengthFourAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 17h2v-6h-2m0 10h2v-2h-2M1 21h16V9h4V1"/>`),
		g.Group(children),
	)
}