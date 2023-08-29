package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintedScreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPaintedScreen0"><g fill="none"><rect width="38" height="28" x="5" y="10" fill="#fff" stroke="#fff" stroke-width="4" rx="3"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13 24v-7"/><rect width="4" height="4" x="15" y="29" fill="#000" rx="2" transform="rotate(90 15 29)"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPaintedScreen0)"/>`),
		g.Group(children),
	)
}