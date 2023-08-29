package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModifyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTModifyTwo0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="40" height="30" x="4" y="9" fill="#555" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 9v30"/><path stroke-linecap="round" d="M20 9h-8m8 30h-8"/><path stroke-linecap="round" stroke-linejoin="round" d="m23 31l14-14m-12 2l-2-2m14 14l-2-2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTModifyTwo0)"/>`),
		g.Group(children),
	)
}