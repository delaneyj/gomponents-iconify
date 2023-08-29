package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartRadial(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 9 9v-1h2v1c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1h1v2h-1Zm0 4a5 5 0 1 0 3.515 1.444l-.711-.703l1.406-1.423l.71.703A7 7 0 1 1 12 5h1v2h-1Zm0 4a1 1 0 1 0 1 1v-1h2v1a3 3 0 1 1-3-3h1v2h-1Z"/>`),
		g.Group(children),
	)
}