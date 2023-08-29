package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 6a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h18a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3H3Zm6 2H7v5a1 1 0 1 1-2 0V8H3a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1h-2v3a1 1 0 1 1-2 0V8h-2v5a1 1 0 1 1-2 0V8h-2v3a1 1 0 1 1-2 0V8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}