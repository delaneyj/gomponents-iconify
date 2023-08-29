package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListUiAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.5 8h-14a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2zm0 5h-10a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2zm0 5h-6a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2z"/><circle cx="3.5" cy="7" r="1" fill="currentColor"/><circle cx="7.5" cy="12" r="1" fill="currentColor"/><circle cx="11.5" cy="17" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}