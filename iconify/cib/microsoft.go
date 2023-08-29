package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microsoft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.198 32H0V16.802h15.198zM32 32H16.802V16.802H32zM15.198 15.198H0V0h15.198zm16.802 0H16.802V0H32z"/>`),
		g.Group(children),
	)
}