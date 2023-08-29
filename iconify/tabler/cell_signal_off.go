package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellSignalOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 20H4.731a.731.731 0 0 1-.517-1.249l7.265-7.264m2-2l5.272-5.272A.731.731 0 0 1 20 4.732v11.269M3 3l18 18"/>`),
		g.Group(children),
	)
}