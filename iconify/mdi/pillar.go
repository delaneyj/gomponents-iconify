package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pillar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 5h12a1 1 0 0 1 1 1a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1a1 1 0 0 1 1-1m15-3v2H3V2h18m-6 6h2v14h-2V8M7 8h2v14H7V8m4 0h2v14h-2V8Z"/>`),
		g.Group(children),
	)
}