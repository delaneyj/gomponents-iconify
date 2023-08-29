package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Key(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 9a7 7 0 1 1 5.562 6.852L12 17.414a2 2 0 0 1-1.414.586H10a2 2 0 0 1-2 2a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-2.586A2 2 0 0 1 2.586 16l5.562-5.562A7.026 7.026 0 0 1 8 9zm7-3a1 1 0 1 0 0 2a1 1 0 0 1 1 1a1 1 0 1 0 2 0a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}