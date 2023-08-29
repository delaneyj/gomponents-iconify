package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.2 8.25L11.85 3h.3l2.65 5.25H9.2Zm2.05 11.85L2.625 9.75h8.625V20.1Zm1.5 0V9.75h8.625L12.75 20.1Zm3.7-11.85L13.85 3h3.925q.575 0 1.038.3t.737.8l2.075 4.15H16.45Zm-14.075 0L4.45 4.1q.275-.5.738-.8T6.225 3h3.925l-2.6 5.25H2.375Z"/>`),
		g.Group(children),
	)
}