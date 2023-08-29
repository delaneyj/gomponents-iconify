package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundDomainVerification(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.23 15.83c.39.39 1.02.39 1.41 0l4.24-4.24c.39-.39.39-1.02 0-1.42a.996.996 0 0 0-1.41 0l-3.54 3.53l-1.41-1.41c-.39-.39-1.02-.39-1.42 0s-.39 1.02 0 1.41l2.13 2.13z"/><path fill="currentColor" d="M19 4H5a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h14c1.1 0 2-.9 2-2V6a2 2 0 0 0-2-2zm0 13c0 .55-.45 1-1 1H6c-.55 0-1-.45-1-1V8h14v9z"/>`),
		g.Group(children),
	)
}