package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChartOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaPieChartOutline0"><g id="evaPieChartOutline1"><g id="evaPieChartOutline2" fill="currentColor"><path d="M13 2a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1a9 9 0 0 0-9-9Zm1 8V4.07A7 7 0 0 1 19.93 10Z"/><path d="M20.82 14.06a1 1 0 0 0-1.28.61A8 8 0 1 1 9.33 4.46a1 1 0 0 0-.66-1.89a10 10 0 1 0 12.76 12.76a1 1 0 0 0-.61-1.27Z"/></g></g></g>`),
		g.Group(children),
	)
}