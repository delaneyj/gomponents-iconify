package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeautyInstrument(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBeautyInstrument0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" d="M17 34h14m-15-7h16"/><rect width="30" height="16" x="9" y="4" fill="#fff" stroke="#fff" rx="4"/><path stroke="#fff" d="m14 20l5 24h10l5-24"/><path stroke="#000" d="M17 12h13"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBeautyInstrument0)"/>`),
		g.Group(children),
	)
}