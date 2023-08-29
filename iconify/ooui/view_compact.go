package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewCompact(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 2h4v4H2zm12 0h4v4h-4zM8 2h4v4H8zM2 14h4v4H2zm12 0h4v4h-4zm-6 0h4v4H8zM2 8h4v4H2zm12 0h4v4h-4zM8 8h4v4H8z"/>`),
		g.Group(children),
	)
}