package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthProducts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHealthProducts0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M33 12H15l-5 5.843v20.249L15 44h18l5-5.908v-20.25L33 12Z"/><path stroke="#000" d="M19 20h10"/><path stroke="#fff" d="M33 12V7a3 3 0 0 0-3-3H18a3 3 0 0 0-3 3v5"/><circle cx="24" cy="32" r="5" stroke="#000"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHealthProducts0)"/>`),
		g.Group(children),
	)
}