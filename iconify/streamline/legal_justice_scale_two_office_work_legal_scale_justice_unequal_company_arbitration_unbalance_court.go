package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LegalJusticeScaleTwoOfficeWorkLegalScaleJusticeUnequalCompanyArbitrationUnbalanceCourt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 11L3 4.82L.5 11a2.5 2.5 0 0 0 5 0Zm8-4L11 1.18L8.5 7a2.5 2.5 0 0 0 5 0Zm-12-1.5l11-5M7 3V.5"/>`),
		g.Group(children),
	)
}