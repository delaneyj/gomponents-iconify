package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingSigmaFormulaTextFormatSigmaFormattingSum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 13.5H1a.5.5 0 0 1-.46-.31a.47.47 0 0 1 .11-.54l5.29-5.3a.5.5 0 0 0 0-.7L.65 1.35A.47.47 0 0 1 .54.81A.5.5 0 0 1 1 .5h12.5"/>`),
		g.Group(children),
	)
}