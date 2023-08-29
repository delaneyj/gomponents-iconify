package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MeasuringCup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMeasuringCup0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M36 6h8l-1.936 14H36"/><path fill="#555" d="M9 6h27v34a2 2 0 0 1-2 2H11a2 2 0 0 1-2-2l-.001-23.5c0-.944-.444-1.828-1.16-2.443C5.148 11.75-.591 6 8.999 6Z"/><path d="M26 15h4m-4 8h4m-4 8h4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMeasuringCup0)"/>`),
		g.Group(children),
	)
}