package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M31 10L14 20"/><path stroke-linecap="round" stroke-linejoin="round" d="M34 17L14 29"/><path stroke-linecap="round" stroke-linejoin="round" d="M34 26L15 38"/><rect width="30" height="6" x="9" y="4" fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round"/><rect width="30" height="6" x="9" y="38" fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round"/><path stroke-linecap="round" stroke-linejoin="round" d="M14 10V38"/><path d="M34 10V38"/></g>`),
		g.Group(children),
	)
}