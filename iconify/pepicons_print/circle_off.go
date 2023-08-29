package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M10.75 15a4.25 4.25 0 1 0 0-8.5a4.25 4.25 0 0 0 0 8.5Zm0 2a6.25 6.25 0 1 0 0-12.5a6.25 6.25 0 0 0 0 12.5Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M10 5a5 5 0 1 0 0 10a5 5 0 0 0 0-10Zm-6 5a6 6 0 1 1 12 0a6 6 0 0 1-12 0Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}