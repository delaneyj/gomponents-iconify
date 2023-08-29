package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21 4v24h8V4h-8zm2 2h4v20h-4V6zM3 10v18h8V10H3zm2 2h4v14H5V12zm7 4v12h8V16h-8zm2 2h4v8h-4v-8z"/>`),
		g.Group(children),
	)
}