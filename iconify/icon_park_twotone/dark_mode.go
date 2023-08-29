package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DarkMode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDarkMode0"><g fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="4"><path d="m24.003 4l5.27 5.27h9.457v9.456l5.27 5.27l-5.27 5.278v9.456h-9.456L24.004 44l-5.278-5.27H9.27v-9.456L4 23.997l5.27-5.27V9.27h9.456L24.003 4Z"/><path d="M27 17c0 8-5 9-10 9c0 4 6.5 8 12 4s2-13-2-13Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDarkMode0)"/>`),
		g.Group(children),
	)
}