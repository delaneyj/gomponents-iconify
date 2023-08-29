package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Potentiometer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><ellipse cx="24" cy="20" fill="#2F88FF" rx="20" ry="8"/><path d="M24 4V19"/><path d="M16 34V42"/><path d="M24 34V44"/><path d="M32 34V42"/><path d="M4 20V32.3636C4 32.3636 5.11012 34.7605 8.5 37.0288"/><path d="M44 20V32.3636C44 32.3636 42.8899 34.7605 39.5 37.0288"/></g>`),
		g.Group(children),
	)
}