package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" stroke-linecap="round" d="M41 44V20H7V44H41Z"/><path stroke="#fff" stroke-linecap="round" d="M24 44V20"/><path stroke="#000" stroke-linecap="round" d="M41 44H7"/><rect width="40" height="8" x="4" y="12" fill="#2F88FF" stroke="#000"/><path stroke="#000" stroke-linecap="round" d="M16 4L24 12L32 4"/></g>`),
		g.Group(children),
	)
}