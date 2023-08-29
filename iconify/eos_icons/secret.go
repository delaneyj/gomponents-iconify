package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Secret(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 9.5A2.492 2.492 0 0 0 13.511 8H10.49a2.436 2.436 0 1 0 .46 1h2.101a2.5 2.5 0 1 0 4.95.5ZM20 5h-3V2a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v3H4v1h16ZM5 17H4a4 4 0 0 0-4 4v3h8l-1-1l2-2Zm15 0h-1l-4 4l2 2l-1 1h8v-3a4 4 0 0 0-4-4Zm-9 3h2v4h-2zm1-2l-1 2h2l-1-2z"/>`),
		g.Group(children),
	)
}