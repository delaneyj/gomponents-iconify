package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpPages(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3v8h5L7 7l4 1V3H3zm5 10H3v8h8v-5l-4 1l1-4zm9 4l-4-1v5h8v-8h-5l1 4zm4-14h-8v5l4-1l-1 4h5V3z"/>`),
		g.Group(children),
	)
}