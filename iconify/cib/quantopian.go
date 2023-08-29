package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quantopian(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 0h28a2 2 0 0 1 2 2v28a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2zm2 4v24h24V4zm3.198 14.401h3.203v6.401H7.198zM12 12h3.198v12.802H12zm4.802 3.198H20v9.604h-3.198zm4.797-8h3.198v17.604h-3.198z"/>`),
		g.Group(children),
	)
}