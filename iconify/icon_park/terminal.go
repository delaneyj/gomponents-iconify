package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="40" height="32" x="4" y="8" fill="#2F88FF" stroke="#000" rx="2"/><path stroke="#fff" stroke-linecap="round" d="M12 18L19 24L12 30"/><path stroke="#fff" stroke-linecap="round" d="M23 32H36"/></g>`),
		g.Group(children),
	)
}