package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FZeroKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" rx="3"/><rect width="8" height="16" x="26" y="16" fill="#2F88FF" stroke="#fff" rx="4"/><path stroke="#fff" d="M21 16H14V32"/><path stroke="#fff" d="M14 24H21"/></g>`),
		g.Group(children),
	)
}