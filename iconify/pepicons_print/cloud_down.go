package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13 4h-1a4.002 4.002 0 0 0-3.874 3H8a4 4 0 1 0 0 8h8a4 4 0 0 0 .899-7.899A4.002 4.002 0 0 0 13 4Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M11 3h-1a4.002 4.002 0 0 0-3.874 3H6a4 4 0 1 0 0 8h8a4 4 0 0 0 .899-7.899A4.002 4.002 0 0 0 11 3ZM6.901 7l.193-.75A3.002 3.002 0 0 1 10 4h1c1.405 0 2.614.975 2.924 2.325l.14.61l.61.141A3.001 3.001 0 0 1 14 13H6a3 3 0 1 1 0-6h.901Z" clip-rule="evenodd"/><path d="M10 10a.5.5 0 0 1 1 0v7.5a.5.5 0 0 1-1 0V10Z"/><path d="M12.688 15.11a.5.5 0 0 1 .624.78l-2.5 2a.5.5 0 1 1-.624-.78l2.5-2Z"/><path d="M7.688 15.89a.5.5 0 0 1 .624-.78l2.5 2a.5.5 0 1 1-.624.78l-2.5-2Z"/></g>`),
		g.Group(children),
	)
}