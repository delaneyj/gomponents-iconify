package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vial(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSVial0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M30 33V4H18v28.968M30 14h-5m5 6h-5"/><path fill="#fff" d="M18 33v4.699C18 41.178 20.686 44 24 44s6-2.821 6-6.301v-4.635L18 33Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSVial0)"/>`),
		g.Group(children),
	)
}