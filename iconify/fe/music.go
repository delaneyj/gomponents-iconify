package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Music(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feMusic0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMusic1" fill="currentColor"><path id="feMusic2" d="M20 4v13a3 3 0 1 1-2-2.83V6h-8v13a3 3 0 1 1-2-2.83V4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2Z"/></g></g>`),
		g.Group(children),
	)
}