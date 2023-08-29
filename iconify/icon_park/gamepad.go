package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gamepad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="40" height="28" x="4" y="13" fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4" rx="14"/><circle cx="31" cy="22" r="2" fill="#fff"/><circle cx="35" cy="27" r="2" fill="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 27H22M12 27H22"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 7V13M24 7V13"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M17 22V32"/></g>`),
		g.Group(children),
	)
}