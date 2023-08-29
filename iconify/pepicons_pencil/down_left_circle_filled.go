package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownLeftCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilDownLeftCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M10.354 19.354a.5.5 0 0 1-.708 0l-4-4a.5.5 0 0 1 0-.708l4-4a.5.5 0 1 1 .708.708L6.707 15l3.647 3.646a.5.5 0 0 1 0 .708Z"/><path d="M15.75 14.5a2.74 2.74 0 0 0 1.943-.81c.516-.52.807-1.226.807-1.963V6a.5.5 0 0 1 1 0v5.727c0 1-.394 1.959-1.097 2.667A3.739 3.739 0 0 1 15.75 15.5H6a.5.5 0 0 1 0-1h9.75Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilDownLeftCircleFilled0)"/></g>`),
		g.Group(children),
	)
}