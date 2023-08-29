package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiskTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="32" height="22" x="4" y="13" stroke="#000" stroke-width="4" rx="2"/><path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 13H14V35H4V13Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36 19H44V29H36"/><circle cx="30" cy="21" r="2" fill="#000"/><circle cx="30" cy="27" r="2" fill="#000"/></g>`),
		g.Group(children),
	)
}