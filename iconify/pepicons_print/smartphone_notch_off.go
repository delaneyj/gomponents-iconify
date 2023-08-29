package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneNotchOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M6 2.5A1.5 1.5 0 0 1 7.5 1h9A1.5 1.5 0 0 1 18 2.5v16a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 6 18.5v-16Z" opacity=".2"/><path fill-rule="evenodd" d="M5.5.5A1.5 1.5 0 0 0 4 2v16a1.5 1.5 0 0 0 1.5 1.5h9A1.5 1.5 0 0 0 16 18V2A1.5 1.5 0 0 0 14.5.5h-9ZM5 2a.5.5 0 0 1 .5-.5H7v1a.5.5 0 0 0 .5.5h5a.5.5 0 0 0 .5-.5v-1h1.5a.5.5 0 0 1 .5.5v16a.5.5 0 0 1-.5.5h-9A.5.5 0 0 1 5 18V2Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}