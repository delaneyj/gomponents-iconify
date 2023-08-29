package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdbOverNetwork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 14.86C8.223 10.504 15.691 7.547 24 7.547s15.777 2.957 19.5 7.313L24 40.453Zm25.55.815h0l1.496-1.389M16.189 14l1.787 1.66"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 13.208a8.605 8.605 0 0 0-8.624 8.624v3.234h17.249v-3.234A8.605 8.605 0 0 0 24 13.208Z"/><circle cx="27.505" cy="18.285" r=".75" fill="currentColor"/><circle cx="20.496" cy="18.285" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}