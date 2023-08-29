package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Poll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22V8h4v14H3m7 0V2h4v20h-4m7 0v-8h4v8h-4Z"/>`),
		g.Group(children),
	)
}