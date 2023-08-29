package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a10 10 0 1 0 4.93 18.69h.12A10 10 0 0 0 12 2Zm1 2.07A8 8 0 0 1 19.93 11H13ZM12 20a8 8 0 0 1-1-15.93V12a1.09 1.09 0 0 0 .07.35v.15l4 6.87A7.81 7.81 0 0 1 12 20Zm4.83-1.64L13.73 13h6.2a8 8 0 0 1-3.1 5.36Z"/>`),
		g.Group(children),
	)
}