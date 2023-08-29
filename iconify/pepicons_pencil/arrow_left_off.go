package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M8.653 14.384a.5.5 0 0 1-.704-.064l-3.333-4a.5.5 0 0 1 .768-.64l3.333 4a.5.5 0 0 1-.064.704Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8.653 5.616a.5.5 0 0 1 .064.704l-3.333 4a.5.5 0 0 1-.768-.64l3.333-4a.5.5 0 0 1 .704-.064Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M5 10a.5.5 0 0 1 .5-.5H15a.5.5 0 0 1 0 1H5.5A.5.5 0 0 1 5 10Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}