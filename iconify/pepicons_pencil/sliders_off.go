package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlidersOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M3 4.75a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5Zm11.38 0a.5.5 0 0 1 .5-.5h1.62a.5.5 0 0 1 0 1h-1.62a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12.75 6.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0 1a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM3 14.75a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5Zm11.38 0a.5.5 0 0 1 .5-.5h1.62a.5.5 0 0 1 0 1h-1.62a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12.75 16.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0 1a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM3 9.75a.5.5 0 0 1 .5-.5h2.13a.5.5 0 0 1 0 1H3.5a.5.5 0 0 1-.5-.5Zm6.5 0a.5.5 0 0 1 .5-.5h6.5a.5.5 0 0 1 0 1H10a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M7.75 11.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0 1a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}