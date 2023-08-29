package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewDay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 8a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V8Zm11-1h6a1 1 0 0 1 1 1v3h-7V7Zm-2 0H5a1 1 0 0 0-1 1v3h7V7Zm-7 6v3a1 1 0 0 0 1 1h6v-4H4Zm9 4h6a1 1 0 0 0 1-1v-3h-7v4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}