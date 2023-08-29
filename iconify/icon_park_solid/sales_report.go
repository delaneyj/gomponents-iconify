package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SalesReport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSalesReport0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M41 14L24 4L7 14v20l17 10l17-10V14Z"/><path stroke="#000" stroke-linecap="round" d="M24 22v8m8-12v12m-16-4v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSalesReport0)"/>`),
		g.Group(children),
	)
}