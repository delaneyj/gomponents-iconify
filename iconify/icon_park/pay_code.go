package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PayCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M44 4H4V44H44V4Z"/><path stroke="#fff" stroke-linecap="round" d="M12 16V32"/><path stroke="#fff" stroke-linecap="round" d="M20 16V32"/><path stroke="#fff" stroke-linecap="round" d="M28 16V32"/><path stroke="#fff" stroke-linecap="round" d="M36 16V32"/></g>`),
		g.Group(children),
	)
}