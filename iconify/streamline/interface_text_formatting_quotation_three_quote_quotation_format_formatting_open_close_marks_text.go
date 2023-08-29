package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingQuotationThreeQuoteQuotationFormatFormattingOpenCloseMarksText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4.08 7.25a3.82 3.82 0 0 1 0-4"/><circle cx="4.45" cy="6.92" r=".5"/><path d="M1.07 7.25a3.82 3.82 0 0 1 0-4"/><circle cx="1.44" cy="6.92" r=".5"/><path d="M9.92 6.75a3.82 3.82 0 0 1 0 4"/><circle cx="9.55" cy="7.08" r=".5"/><path d="M12.93 6.75a3.82 3.82 0 0 1 0 4"/><circle cx="12.56" cy="7.08" r=".5"/></g>`),
		g.Group(children),
	)
}