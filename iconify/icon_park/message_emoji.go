package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageEmoji(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M44 7H4V37H11V42L21 37H44V7Z"/><path stroke="#fff" d="M31 16V17"/><path stroke="#fff" d="M17 16V17"/><path stroke="#fff" d="M31 25C31 25 29 29 24 29C19 29 17 25 17 25"/></g>`),
		g.Group(children),
	)
}