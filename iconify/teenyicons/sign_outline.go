package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4.5 6.5V6a.5.5 0 0 0-.5.5h.5Zm10 0h.5a.5.5 0 0 0-.5-.5v.5Zm0 6v.5a.5.5 0 0 0 .5-.5h-.5Zm-10 0H4a.5.5 0 0 0 .5.5v-.5ZM1 1v14h1V1H1ZM0 4h15V3H0v1Zm4.5 3h10V6h-10v1Zm9.5-.5v6h1v-6h-1Zm.5 5.5h-10v1h10v-1Zm-9.5.5v-6H4v6h1Zm1-9v3h1v-3H6Zm6 0v3h1v-3h-1Z"/>`),
		g.Group(children),
	)
}