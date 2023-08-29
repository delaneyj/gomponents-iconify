package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurveAdjustment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCurveAdjustment0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M44 4H4v40h40V4Z"/><path stroke-linecap="round" d="M38 10c-6 0-11 4-14 14s-8 14-14 14"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCurveAdjustment0)"/>`),
		g.Group(children),
	)
}