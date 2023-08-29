package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12zm6.498-4.865a1 1 0 0 1 .998-.003L15 10.848V8a1 1 0 1 1 2 0v8a1 1 0 1 1-2 0v-2.848l-6.504 3.716A1 1 0 0 1 7 16V8a1 1 0 0 1 .498-.865z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}