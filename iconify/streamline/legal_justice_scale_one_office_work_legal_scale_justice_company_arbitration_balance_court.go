package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LegalJusticeScaleOneOfficeWorkLegalScaleJusticeCompanyArbitrationBalanceCourt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 9.5L3 4L.5 9.5a2.5 2.5 0 0 0 5 0Zm8 0L11 4L8.5 9.5a2.5 2.5 0 0 0 5 0ZM1.5 4h11M7 4V2"/>`),
		g.Group(children),
	)
}