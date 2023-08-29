package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="36" height="28" x="6" y="6" fill="#2F88FF" rx="3"/><path stroke-linecap="round" d="M14 42L34 42"/><path stroke-linecap="round" d="M24 34V42"/></g>`),
		g.Group(children),
	)
}