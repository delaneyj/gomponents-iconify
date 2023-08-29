package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewGridDetail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSViewGridDetail0"><g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" rx="3"/><path fill="#000" stroke="#000" d="M13 13h8v8h-8z"/><path stroke="#000" stroke-linecap="round" d="M27 13h8m-8 7h8m-22 8h22m-22 7h22"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSViewGridDetail0)"/>`),
		g.Group(children),
	)
}