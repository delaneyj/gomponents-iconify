package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingQuotationOneQuoteQuotationFormatFormattingOpenCloseMarksText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.07 7a3.82 3.82 0 0 1 0-4m3.01 4a3.82 3.82 0 0 1 0-4m8.85 8a3.82 3.82 0 0 0 0-4m-3.01 4a3.82 3.82 0 0 0 0-4"/>`),
		g.Group(children),
	)
}