package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cancer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><circle cx="37" cy="17" r="6" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 13C6 13 12 5 22 5C32 5 38 11 38 11"/><circle cx="11" cy="31" r="6" fill="#2F88FF" transform="rotate(-180 11 31)"/><path stroke-linecap="round" stroke-linejoin="round" d="M42 35C42 35 36 43 26 43C16 43 10 37 10 37"/></g>`),
		g.Group(children),
	)
}