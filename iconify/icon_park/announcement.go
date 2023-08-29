package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Announcement(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="40" height="26" x="4" y="15" fill="#2F88FF" stroke="#000" rx="2"/><path fill="#2F88FF" stroke="#000" stroke-linecap="round" d="M24 7L16 15H32L24 7Z"/><path stroke="#fff" stroke-linecap="round" d="M12 24H30"/><path stroke="#fff" stroke-linecap="round" d="M12 32H20"/></g>`),
		g.Group(children),
	)
}