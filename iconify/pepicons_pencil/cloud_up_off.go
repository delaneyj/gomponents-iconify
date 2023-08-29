package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudUpOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M11 3h-1a4.002 4.002 0 0 0-3.874 3H6a4 4 0 1 0 0 8h8a4 4 0 0 0 .899-7.899A4.002 4.002 0 0 0 11 3ZM6.901 7l.193-.75A3.002 3.002 0 0 1 10 4h1c1.405 0 2.614.975 2.924 2.325l.14.61l.61.141A3.001 3.001 0 0 1 14 13H6a3 3 0 1 1 0-6h.901Z" clip-rule="evenodd"/><path d="M11 16.5a.5.5 0 0 1-1 0V9a.5.5 0 0 1 1 0v7.5Z"/><path d="M8.312 11.39a.5.5 0 0 1-.624-.78l2.5-2a.5.5 0 0 1 .624.78l-2.5 2Z"/><path d="M13.312 10.61a.5.5 0 0 1-.624.78l-2.5-2a.5.5 0 1 1 .624-.78l2.5 2ZM1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}