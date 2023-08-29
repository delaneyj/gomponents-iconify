package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 17c0-2.038-1.188-3.836-3-4.92V5h.5a1.5 1.5 0 0 0 0-3h-9a1.5 1.5 0 0 0 0 3H8v7.08C6.188 13.164 5 14.962 5 17h6v4c0 .55.45 1 1 1s1-.45 1-1v-4h6z"/>`),
		g.Group(children),
	)
}