package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7 9V2a7 7 0 1 0 6.262 3.869L7 9zm7.262-5.131A6.999 6.999 0 0 0 8 0v7l6.262-3.131z"/>`),
		g.Group(children),
	)
}