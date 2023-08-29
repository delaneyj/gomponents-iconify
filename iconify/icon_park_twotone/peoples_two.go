package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeoplesTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPeoplesTwo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M24 20a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/><path d="M12 7.255A6.992 6.992 0 0 0 9 13a6.996 6.996 0 0 0 3.392 6M36 7.255A6.992 6.992 0 0 1 39 13a6.992 6.992 0 0 1-3 5.745"/><path fill="#555" d="M12 40v2h24v-2c0-3.727 0-5.591-.609-7.062a8 8 0 0 0-4.33-4.329C29.591 28 27.727 28 24 28s-5.591 0-7.061.609a8 8 0 0 0-4.33 4.33C12 34.408 12 36.273 12 40Z"/><path d="M44 42v-1.2c0-4.48 0-6.72-.872-8.432a8 8 0 0 0-3.496-3.496M4 42v-1.2c0-4.48 0-6.72.872-8.432a8 8 0 0 1 3.496-3.496"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPeoplesTwo0)"/>`),
		g.Group(children),
	)
}