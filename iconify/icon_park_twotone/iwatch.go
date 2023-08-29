package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iwatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTIwatch0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M36 12H10v24h26V12Z"/><path fill="#555" stroke-linejoin="round" d="M27 24a4 4 0 1 0-8 0a4 4 0 0 0 8 0Z"/><path stroke-linecap="round" d="M15 12c0-.073.008-.146.025-.217l1.607-7A1.026 1.026 0 0 1 17.641 4h10.743c.484 0 .903.326 1.008.783l1.608 7M15 36c0 .073.008.146.025.217l1.607 7c.105.458.524.783 1.009.783h10.743c.484 0 .903-.325 1.008-.783l1.608-7M39 16v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTIwatch0)"/>`),
		g.Group(children),
	)
}