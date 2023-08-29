package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentLayoutCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 8h-2a1 1 0 0 1 0-2h2a1 1 0 0 1 0 2zm0 4h-2a1 1 0 0 1 0-2h2a1 1 0 0 1 0 2zM5 8H3a1 1 0 0 1 0-2h2a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h2a1 1 0 0 1 0 2z" opacity=".5"/><rect width="8" height="8" x="8" y="4" fill="currentColor" rx="1"/><path fill="currentColor" d="M21 16H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm-8 4H3a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2z" opacity=".5"/>`),
		g.Group(children),
	)
}