package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServicePlanOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 2h8v2h-8zm-4 0h2v2h-2zm4 4h8v2h-8zm-4 0h2v2h-2zm4 4h8v2h-8zm-4 0h2v2h-2z"/><circle cx="5" cy="19" r="1" fill="currentColor"/><path fill="currentColor" d="M4 4h2v9H4z"/><path fill="currentColor" d="M7 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1Zm0 19H3V3h4Z"/>`),
		g.Group(children),
	)
}