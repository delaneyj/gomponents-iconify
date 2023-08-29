package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anchor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22a8.01 8.01 0 0 1-8-8a1 1 0 0 1 2 0a6 6 0 0 0 12 0a1 1 0 0 1 2 0a8.01 8.01 0 0 1-8 8Z"/><path fill="currentColor" d="M12 22a1 1 0 0 1-1-1V7a1 1 0 0 1 2 0v14a1 1 0 0 1-1 1zm-5-7H5a1 1 0 0 1 0-2h2a1 1 0 0 1 0 2zm12 0h-2a1 1 0 0 1 0-2h2a1 1 0 0 1 0 2z"/><path fill="currentColor" d="M14 11h-4a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2zm-2-3a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3zm0-4a1 1 0 1 0 1 1a1.001 1.001 0 0 0-1-1z"/>`),
		g.Group(children),
	)
}