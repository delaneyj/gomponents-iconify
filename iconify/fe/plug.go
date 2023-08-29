package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="fePlug0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePlug1" fill="currentColor"><path id="fePlug2" d="M17 11v2h3a1 1 0 0 1 0 2h-3.268A2 2 0 0 1 15 16h-3a4.002 4.002 0 0 1-3.876-3.008A1.01 1.01 0 0 1 8 13H4a1 1 0 0 1 0-2h4c.042 0 .083.003.124.008A4.002 4.002 0 0 1 12 8h3a2 2 0 0 1 1.732 1H20a1 1 0 0 1 0 2h-3Z"/></g></g>`),
		g.Group(children),
	)
}