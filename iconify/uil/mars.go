package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mars(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.42 4.12a1 1 0 0 0-.54-.54a1 1 0 0 0-.38-.08h-4a1 1 0 0 0 0 2h1.59l-2.4 2.4a7 7 0 1 0 1.41 1.41l2.4-2.4V8.5a1 1 0 0 0 2 0v-4a1 1 0 0 0-.08-.38ZM14 17a5 5 0 1 1 0-7a5 5 0 0 1 0 7Z"/>`),
		g.Group(children),
	)
}