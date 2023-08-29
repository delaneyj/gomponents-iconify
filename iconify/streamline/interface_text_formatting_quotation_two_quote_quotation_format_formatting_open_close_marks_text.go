package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingQuotationTwoQuoteQuotationFormatFormattingOpenCloseMarksText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.18 5.44a1 1 0 0 1 .42.81a.9.9 0 0 1-1.1 1C.54 7 .09 5.57 1 3.25m4.7 2.19a1 1 0 0 1 .43.81a.9.9 0 0 1-1.1 1c-1-.19-1.41-1.66-.53-4m7.32 5.31a1 1 0 0 1-.42-.81a.9.9 0 0 1 1.1-1c1 .19 1.41 1.66.52 4M8.3 8.56a1 1 0 0 1-.43-.81a.9.9 0 0 1 1.1-1c1 .19 1.41 1.66.53 4"/>`),
		g.Group(children),
	)
}