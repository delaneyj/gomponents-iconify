package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cinema(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M14 7.5v2a.5.5 0 0 1-1 0s.06-.5-1-.5h-1v2.5a.5.5 0 0 1-.5.5h-8a.5.5 0 0 1-.5-.5v-4a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 .5.5V8h1c1.06 0 1-.5 1-.5a.5.5 0 0 1 1 0zM4 3a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm0 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm4.5-4a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5zm0 4a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3z"/>`),
		g.Group(children),
	)
}