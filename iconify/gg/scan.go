package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 3h2v18h-2V3ZM5 8a1 1 0 0 1 1-1h3V5H6a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h3v-2H6a1 1 0 0 1-1-1V8Zm14 0a1 1 0 0 0-1-1h-3V5h3a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3h-3v-2h3a1 1 0 0 0 1-1V8Z"/>`),
		g.Group(children),
	)
}