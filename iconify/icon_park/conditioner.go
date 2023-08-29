package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Conditioner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="32" height="20" x="8" y="24" fill="#2F88FF" stroke="#000"/><path stroke="#000" d="M37 24V17H11V24"/><path stroke="#000" d="M17 17C19 14.8333 20 10.5 20 4C23 4 30 9.41667 30 16.7327"/><rect width="16" height="6" x="16" y="31" stroke="#fff"/></g>`),
		g.Group(children),
	)
}