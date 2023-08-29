package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func More(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-2a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm7 2a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-2a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm10-1a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm-2 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}