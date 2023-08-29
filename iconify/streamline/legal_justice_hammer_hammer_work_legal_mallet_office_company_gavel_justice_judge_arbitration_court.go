package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LegalJusticeHammerHammerWorkLegalMalletOfficeCompanyGavelJusticeJudgeArbitrationCourt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 13.29H8m-1 0v-2.5H1.5v2.5"/><rect width="7.07" height="4.24" x="3.96" y="2.17" rx="1" transform="rotate(-45 7.499 4.294)"/><path d="m9 5.79l4.5 4.5"/></g>`),
		g.Group(children),
	)
}