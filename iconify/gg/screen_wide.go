package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenWide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 16H3a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3h18a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-8v1h2a1 1 0 1 1 0 2H9a1 1 0 1 1 0-2h2v-1ZM3 7h18a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}