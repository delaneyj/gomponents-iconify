package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M0 22h24M22 2h-4v16h4V2ZM6 6H2v12h4V6Zm8 12h-4v-8h4v8Z"/>`),
		g.Group(children),
	)
}