package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Potentiometer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><ellipse cx="24" cy="20" rx="20" ry="8"/><path d="M24 4v15m-8 15v8m8-8v10m8-10v8M4 20v12.364s1.11 2.397 4.5 4.665M44 20v12.364s-1.11 2.397-4.5 4.665"/></g>`),
		g.Group(children),
	)
}