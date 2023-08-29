package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Die(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 1a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2zm2 16a2 2 0 1 1 2-2a2 2 0 0 1-2 2zM5 7a2 2 0 1 1 2-2a2 2 0 0 1-2 2zm5 5a2 2 0 1 1 2-2a2 2 0 0 1-2 2zm5 5a2 2 0 1 1 2-2a2 2 0 0 1-2 2zm0-10a2 2 0 1 1 2-2a2 2 0 0 1-2 2z"/>`),
		g.Group(children),
	)
}