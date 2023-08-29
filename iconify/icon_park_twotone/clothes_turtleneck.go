package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesTurtleneck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTClothesTurtleneck0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M37 23v14m-26 0v7h26v-7m-26 0H4V23c0-3 2-6.5 5-9s9-4 9-4h12s6 1.5 9 4s5 6 5 9v14h-7m-26 0V23"/><path fill="#555" d="M30 10H18V4h12v6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTClothesTurtleneck0)"/>`),
		g.Group(children),
	)
}