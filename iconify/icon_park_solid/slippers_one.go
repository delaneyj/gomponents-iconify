package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlippersOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSlippersOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 29h40v6H4v-6Z"/><path fill="#fff" d="M7 22c-3 4-3 7-3 7h26v-8c0-2.5-1.5-5.5-5-6s-13.124.5-18 7Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSlippersOne0)"/>`),
		g.Group(children),
	)
}