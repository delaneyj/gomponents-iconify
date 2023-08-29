package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Table(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="32" height="40" x="8" y="4" fill="#2F88FF" stroke="#000" rx="2"/><path stroke="#fff" d="M14 16H34"/><path stroke="#fff" d="M14 24H34"/><path stroke="#fff" d="M14 32H34"/><path stroke="#fff" d="M18 12V36"/></g>`),
		g.Group(children),
	)
}