package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingQuotationFourQuoteQuotationFormatFormattingOpenCloseMarksText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 4L1 7m2-3l.5 3m7.75 0l-.5 3m2.75-3l-.5 3"/>`),
		g.Group(children),
	)
}