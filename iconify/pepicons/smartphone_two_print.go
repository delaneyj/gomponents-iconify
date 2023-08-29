package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneTwoPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path d="M7.636 1h8.728C17.267 1 18 1.672 18 2.5v16c0 .828-.733 1.5-1.636 1.5H7.636C6.733 20 6 19.328 6 18.5v-16C6 1.672 6.733 1 7.636 1Z" opacity=".8"/><path fill-rule="evenodd" d="M14.5.5h-9A1.5 1.5 0 0 0 4 2v16a1.5 1.5 0 0 0 1.5 1.5h9A1.5 1.5 0 0 0 16 18V2A1.5 1.5 0 0 0 14.5.5ZM5 2a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 .5.5v16a.5.5 0 0 1-.5.5h-9A.5.5 0 0 1 5 18V2Z" clip-rule="evenodd"/><path d="M10 18a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/></g>`),
		g.Group(children),
	)
}