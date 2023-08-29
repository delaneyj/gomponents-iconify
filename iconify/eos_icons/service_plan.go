package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServicePlan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1ZM5 21a2 2 0 1 1 2-2a2 2 0 0 1-2 2Zm2-9H3V3h4Z"/><circle cx="5.01" cy="19.01" r="1" fill="currentColor"/><path fill="currentColor" d="M14 2h8v2.02h-8zm-4 0h2.01v2.02H10zm4 4h8v2.02h-8zm-4 0h2.01v2.02H10zm4 4h8v2.02h-8zm-4 0h2.01v2.02H10z"/>`),
		g.Group(children),
	)
}