package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gamepad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSGamepad0"><g fill="none"><rect width="40" height="28" x="4" y="13" fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" rx="14"/><circle cx="31" cy="22" r="2" fill="#000"/><circle cx="35" cy="27" r="2" fill="#000"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 27h10m-10 0h10"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 7v6m0-6v6"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M17 22v10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSGamepad0)"/>`),
		g.Group(children),
	)
}