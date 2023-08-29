package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineChartUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 21H4a1 1 0 0 1-1-1V3h2v16h16v2ZM8.373 16L7 14.656L11.856 9.9a.985.985 0 0 1 1.373 0l2.227 2.181L19.627 8L21 9.344L16.144 14.1a.985.985 0 0 1-1.373 0l-2.228-2.182L8.374 16h-.001Z"/>`),
		g.Group(children),
	)
}