package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M5 8a1.5 1.5 0 0 1 1.5-1.5h10a1.5 1.5 0 0 1 0 3h-10A1.5 1.5 0 0 1 5 8Zm0 6a1.5 1.5 0 0 1 1.5-1.5h10a1.5 1.5 0 0 1 0 3h-10A1.5 1.5 0 0 1 5 14Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M4.5 7a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H5a.5.5 0 0 1-.5-.5Zm0 6a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H5a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}