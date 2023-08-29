package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Down(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m192 284l186-162q15-17 2-30q-17-17-30-2L192 228L36 90Q23 75 6 92q-14 14 3 30z"/>`),
		g.Group(children),
	)
}