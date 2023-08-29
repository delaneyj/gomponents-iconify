package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snaptube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="39" height="27" x="4.5" y="10.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3" ry="3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.514 25.301l7.5 7.199l7.472-7.199h-4.461v-3.177h-6.016v3.177h-4.495Zm4.495-7.712h6.016v2.249h-6.016zm0-2.089h6.016"/>`),
		g.Group(children),
	)
}