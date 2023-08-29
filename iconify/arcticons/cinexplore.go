package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cinexplore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.907 30.149L7.1 10.631m-.007 26.698l33.814-19.522M17.871 4.5v38.997m23.447-14.884v-9.27a3.073 3.073 0 0 0-1.536-2.66L19.395 4.912a3.073 3.073 0 0 0-3.054-.01L8.236 9.505a3.072 3.072 0 0 0-1.554 2.672v23.614a3.073 3.073 0 0 0 1.536 2.661l8.028 4.635a3.073 3.073 0 0 0 3.073 0l20.463-11.814a3.072 3.072 0 0 0 1.536-2.66Z"/>`),
		g.Group(children),
	)
}