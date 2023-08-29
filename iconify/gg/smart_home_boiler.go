package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartHomeBoiler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 5a4 4 0 0 1 4-4h6a4 4 0 0 1 4 4v16h-3.856l.742 2h-2l-.742-2h-2l.742 2h-2l-.742-2H5V5Zm4-2h6a2 2 0 0 1 2 2v2H7V5a2 2 0 0 1 2-2ZM7 9h10v10H7V9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}