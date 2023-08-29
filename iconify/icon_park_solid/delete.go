package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Delete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDelete0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M9 10v34h30V10H9Z"/><path stroke="#000" stroke-linecap="round" d="M20 20v13m8-13v13"/><path stroke="#fff" stroke-linecap="round" d="M4 10h40"/><path fill="#fff" stroke="#fff" d="m16 10l3.289-6h9.488L32 10H16Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDelete0)"/>`),
		g.Group(children),
	)
}