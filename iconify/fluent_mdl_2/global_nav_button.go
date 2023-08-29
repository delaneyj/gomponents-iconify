package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobalNavButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 640H0V512h2048v128zm0 1024H0v-128h2048v128zm0-513H0v-127h2048v127z"/>`),
		g.Group(children),
	)
}