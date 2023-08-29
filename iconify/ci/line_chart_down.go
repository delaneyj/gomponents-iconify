package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineChartDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 21H4a1 1 0 0 1-1-1V3h2v16h16v2Zm-1.373-5l-4.17-4.082l-2.228 2.182a.985.985 0 0 1-1.373 0L7 9.344L8.373 8l4.17 4.082L14.77 9.9a.985.985 0 0 1 1.373 0L21 14.656L19.627 16Z"/>`),
		g.Group(children),
	)
}