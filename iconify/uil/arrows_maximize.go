package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsMaximize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 14a1 1 0 0 0-1 1v3.59L5.41 4H9a1 1 0 0 0 0-2H3a1 1 0 0 0-.38.08a1 1 0 0 0-.54.54A1 1 0 0 0 2 3v6a1 1 0 0 0 2 0V5.41L18.59 20H15a1 1 0 0 0 0 2h6a1 1 0 0 0 .38-.08a1 1 0 0 0 .54-.54A1 1 0 0 0 22 21v-6a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}