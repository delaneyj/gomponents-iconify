package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M25.5 2.5h-10c-2.53 0-3 .529-3 3v15H.5l20 20l20-20h-12v-15c0-2.439-.5-3-3-3z"/>`),
		g.Group(children),
	)
}