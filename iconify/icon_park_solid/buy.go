package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBuy0"><g fill="none" stroke-width="4"><path fill="#fff" fill-rule="evenodd" stroke="#fff" stroke-linejoin="round" d="M6 15h36l-2 27H8L6 15Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M16 19V6h16v13"/><path stroke="#000" stroke-linecap="round" d="M16 34h16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBuy0)"/>`),
		g.Group(children),
	)
}