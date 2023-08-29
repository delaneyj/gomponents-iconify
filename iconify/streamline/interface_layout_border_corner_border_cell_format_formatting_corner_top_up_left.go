package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceLayoutBorderCornerBorderCellFormatFormattingCornerTopUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 11.5v1a1 1 0 0 1-1 1h-1m2-9v4m-5 5h-3m-5-1v-11a1 1 0 0 1 1-1h11"/>`),
		g.Group(children),
	)
}