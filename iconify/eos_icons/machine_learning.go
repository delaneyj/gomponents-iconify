package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MachineLearning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 8h2v1h-2zm0-4h2v1h-2zm0 6h2v1h-2z"/><path fill="currentColor" d="M21 12V9a13.124 13.124 0 0 0-8.354 3h-1.292A13.124 13.124 0 0 0 3 9v3a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1v4a13.153 13.153 0 0 1 9 3.55A13.2 13.2 0 0 1 21 20v-4a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1Z"/><circle cx="9" cy="4" r="1" fill="currentColor"/><circle cx="15" cy="4" r="1" fill="currentColor"/><path fill="currentColor" d="M16 8H8a3.003 3.003 0 0 1-3-3V3a3.003 3.003 0 0 1 3-3h8a3.003 3.003 0 0 1 3 3v2a3.003 3.003 0 0 1-3 3ZM8 2a1.001 1.001 0 0 0-1 1v2a1.001 1.001 0 0 0 1 1h8a1.001 1.001 0 0 0 1-1V3a1.001 1.001 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}