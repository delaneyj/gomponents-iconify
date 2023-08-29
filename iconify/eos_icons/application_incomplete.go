package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationIncomplete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3v2h22V3a2 2 0 0 0-2-2H3a2 2 0 0 0-2 2Zm3 1a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm3 0a1 1 0 1 1 1-1a1.003 1.003 0 0 1-1 1ZM1 6v14a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V6Zm11 13a5 5 0 0 1 0-10v5h5a5 5 0 0 1-5 5Z"/>`),
		g.Group(children),
	)
}