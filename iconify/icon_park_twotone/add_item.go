package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddItem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAddItem0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 30v-6"/><path fill="#555" d="M6 30h36v12H6V30Z"/><path d="M6 12.5v-1M6 18v-1M6 7V6m36 6.5v-1m0 6.5v-1m0-10V6m0 12h-1M7 18H6M7 6H6m8 0h-1m8 0h-1m1 12h-1m8-12h-1M14 18h-1m15 0h-1m8-12h-1m1 12h-1m8-12h-1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAddItem0)"/>`),
		g.Group(children),
	)
}