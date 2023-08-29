package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagNone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23.092 10.322L30 3.414L28.586 2L2 28.586L3.414 30l6.908-6.908l6.322 6.323a2.001 2.001 0 0 0 2.828 0l9.943-9.943a2.001 2.001 0 0 0 0-2.828zM18.058 28l-6.322-6.322l9.942-9.942L28 18.058zM10 14a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4zm0-6a2 2 0 1 0 1.998 2.004A2.002 2.002 0 0 0 10 8z"/><path fill="currentColor" d="m7.493 20.263l1.414-1.414L4 13.94V4h9.942l4.907 4.907l1.414-1.414l-4.908-4.907A2 2 0 0 0 13.941 2H4a2 2 0 0 0-2 2v9.941a2 2 0 0 0 .586 1.414Z"/>`),
		g.Group(children),
	)
}