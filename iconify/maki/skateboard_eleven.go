package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkateboardEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M8 6H3a2.002 2.002 0 0 1-2-2h1a1.001 1.001 0 0 0 1 1h5a1.001 1.001 0 0 0 1-1h1a2.002 2.002 0 0 1-2 2zm0 1a1 1 0 1 0 1 1a1 1 0 0 0-1-1zM3 7a1 1 0 1 0 1 1a1 1 0 0 0-1-1z" fill="currentColor"/>`),
		g.Group(children),
	)
}