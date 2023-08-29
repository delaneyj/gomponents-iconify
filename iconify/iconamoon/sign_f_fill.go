package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignFFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.295 4c.39 0 .705.316.705.705V5a1 1 0 1 0 2 0v-.295A2.705 2.705 0 0 0 15.295 2a2.73 2.73 0 0 0-2.62 1.96L11.487 8H10.5a1 1 0 0 0 0 2h.399l-2.301 7.823A1.64 1.64 0 0 1 7.025 19H7a1 1 0 1 0 0 2h.025a3.64 3.64 0 0 0 3.492-2.613L12.983 10H14a1 1 0 1 0 0-2h-.428l1.022-3.475A.73.73 0 0 1 15.295 4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}