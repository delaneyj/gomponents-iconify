package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SleepTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M4 12L4 36"/><path stroke-linecap="round" stroke-linejoin="round" d="M44 29L44 36"/><path stroke-linecap="round" stroke-linejoin="round" d="M44 29L4 29"/><path fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round" d="M22 16L22 29L44 29L44 19C44 17.3431 42.6569 16 41 16L22 16Z"/><circle cx="13" cy="20" r="3" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}