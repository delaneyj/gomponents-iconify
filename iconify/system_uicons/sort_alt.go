package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.5 8.5l4 4l4-4m-4-6v10m-4 .044L6.5 8.5l-4 4.044m4-4.044v10"/>`),
		g.Group(children),
	)
}