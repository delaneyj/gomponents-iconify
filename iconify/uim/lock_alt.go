package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 7a3 3 0 1 1 6 0v2h2V7A5 5 0 0 0 7 7v2h2V7zm3 11a1 1 0 0 1-1-1v-3a1 1 0 1 1 2 0v3a1 1 0 0 1-1 1z" opacity=".5"/><path fill="currentColor" d="M17 9H7a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3zm-4 8a1 1 0 0 1-2 0v-3a1 1 0 1 1 2 0v3z"/>`),
		g.Group(children),
	)
}