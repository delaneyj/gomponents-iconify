package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneCutoutOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6 2.5A1.5 1.5 0 0 1 7.5 1h8A1.5 1.5 0 0 1 17 2.5v16a1.5 1.5 0 0 1-1.5 1.5h-8A1.5 1.5 0 0 1 6 18.5v-16Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M4 2A1.5 1.5 0 0 1 5.5.5h9A1.5 1.5 0 0 1 16 2v16a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 4 18V2Zm1.5-.5A.5.5 0 0 0 5 2v16a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5V2a.5.5 0 0 0-.5-.5h-9Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M7.5 3a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2h-3a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}