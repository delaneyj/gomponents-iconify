package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudUpPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13 4h-1a4.002 4.002 0 0 0-3.874 3H8a4 4 0 1 0 0 8h8a4 4 0 0 0 .899-7.899A4.002 4.002 0 0 0 13 4Z" clip-rule="evenodd" opacity=".8"/><path fill-rule="evenodd" d="M10 3H9a4.002 4.002 0 0 0-3.874 3H5a4 4 0 1 0 0 8h8a4 4 0 0 0 .899-7.899A4.002 4.002 0 0 0 10 3ZM5.901 7l.193-.75A3.002 3.002 0 0 1 9 4h1c1.405 0 2.614.975 2.924 2.325l.14.61l.61.141A3.001 3.001 0 0 1 13 13H5a3 3 0 1 1 0-6h.901Z" clip-rule="evenodd"/><path d="M10 16.5a.5.5 0 0 1-1 0V9a.5.5 0 0 1 1 0v7.5Z"/><path d="M7.312 11.39a.5.5 0 0 1-.624-.78l2.5-2a.5.5 0 0 1 .624.78l-2.5 2Z"/><path d="M12.312 10.61a.5.5 0 0 1-.624.78l-2.5-2a.5.5 0 0 1 .624-.78l2.5 2Z"/></g>`),
		g.Group(children),
	)
}