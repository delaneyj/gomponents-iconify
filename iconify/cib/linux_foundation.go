package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinuxFoundation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6.401 12.802v12.797h12.797V32H0V12.802zM32 0v32h-9.599v-6.401h3.198V6.401H6.401v3.198H0V0z"/>`),
		g.Group(children),
	)
}