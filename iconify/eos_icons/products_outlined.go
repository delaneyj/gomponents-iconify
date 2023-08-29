package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProductsOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="5" cy="19" r="1" fill="currentColor"/><path fill="currentColor" d="M4 4h2v9H4z"/><path fill="currentColor" d="M7 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1Zm0 19H3V3h4Z"/><circle cx="12" cy="19" r="1" fill="currentColor"/><path fill="currentColor" d="M11 4h2v9h-2z"/><path fill="currentColor" d="M14 2h-4a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1Zm0 19h-4V3h4Z"/><circle cx="19" cy="19" r="1" fill="currentColor"/><path fill="currentColor" d="M18 4h2v9h-2z"/><path fill="currentColor" d="M21 2h-4a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1Zm0 19h-4V3h4Z"/>`),
		g.Group(children),
	)
}