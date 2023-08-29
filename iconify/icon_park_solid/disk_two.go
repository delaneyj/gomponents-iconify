package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiskTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDiskTwo0"><g fill="none"><rect width="32" height="22" x="4" y="13" stroke="#fff" stroke-width="4" rx="2"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 13h10v22H4V13Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36 19h8v10h-8"/><circle cx="30" cy="21" r="2" fill="#fff"/><circle cx="30" cy="27" r="2" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDiskTwo0)"/>`),
		g.Group(children),
	)
}