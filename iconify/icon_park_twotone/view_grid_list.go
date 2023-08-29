package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewGridList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTViewGridList0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#555" rx="3"/><path fill="#555" d="M13 13h8v8h-8zm0 14h8v8h-8z"/><path stroke-linecap="round" d="M27 28h8m-8 7h8m-8-22h8m-8 7h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTViewGridList0)"/>`),
		g.Group(children),
	)
}