package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextMessage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M44 7H4V37H19L24 42L29 37H44V7Z"/><path stroke="#fff" d="M14 16H20"/><path stroke="#fff" d="M14 24H16"/><path stroke="#fff" d="M29 14L36 28"/><path stroke="#fff" d="M28.9998 13.9998L21.9998 27.9998"/><path stroke="#fff" d="M24 24H34"/></g>`),
		g.Group(children),
	)
}