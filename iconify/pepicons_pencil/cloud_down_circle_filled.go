package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudDownCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilCloudDownCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M14 6h-1a4.002 4.002 0 0 0-3.874 3H9a4 4 0 1 0 0 8h8a4 4 0 0 0 .899-7.899A4.002 4.002 0 0 0 14 6Zm-4.099 4l.193-.75A3.002 3.002 0 0 1 13 7h1c1.405 0 2.614.975 2.924 2.325l.14.61l.61.141A3.001 3.001 0 0 1 17 16H9a3 3 0 1 1 0-6h.901Z" clip-rule="evenodd"/><path d="M13 13a.5.5 0 0 1 1 0v7.5a.5.5 0 0 1-1 0V13Z"/><path d="M15.688 18.11a.5.5 0 0 1 .624.78l-2.5 2a.5.5 0 1 1-.624-.78l2.5-2Z"/><path d="M10.688 18.89a.5.5 0 0 1 .624-.78l2.5 2a.5.5 0 1 1-.624.78l-2.5-2Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilCloudDownCircleFilled0)"/></g>`),
		g.Group(children),
	)
}