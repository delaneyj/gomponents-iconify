package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculationOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 1h22v22H1V1Zm2 2v18h18V3H3Zm6 2v2h2v2H9v2H7V9H5V7h2V5h2Zm4 2h6v2h-6V7Zm-6.414 6.172L8 14.586l1.414-1.414l1.414 1.414L9.414 16l1.415 1.414l-1.415 1.415L8 17.414L6.586 18.83L5.17 17.414L6.586 16l-1.414-1.414l1.414-1.414ZM13 13.5h6v2h-6v-2Zm0 3h6v2h-6v-2Z"/>`),
		g.Group(children),
	)
}