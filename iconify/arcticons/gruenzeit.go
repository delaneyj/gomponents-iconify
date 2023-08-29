package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gruenzeit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.072 44.5s-12.424-32.303 0-32.303s0 32.303 0 32.303Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.102 19.651s1.238.207 1.864.249c.62.041 1.863 0 1.863 0m-3.727 3.479a12.07 12.07 0 0 0 1.242.248a11.84 11.84 0 0 0 1.243 0m7.454-1.491s-1.238.207-1.863.249c-.62.04-1.864 0-1.864 0m3.727 2.236a12.073 12.073 0 0 1-1.242.248a11.81 11.81 0 0 1-1.242 0m1.243 7.207a12.063 12.063 0 0 1-1.243.248a11.84 11.84 0 0 1-1.243 0m-3.727 4.721a12.063 12.063 0 0 0 1.242.249a11.839 11.839 0 0 0 1.243 0m3.901-23.32s5.86-6.175 3.553-7.99c-1.363-1.253-4.97 2.486-4.97 2.486s-.632-4.97-2.484-4.97s-2.485 4.97-2.485 4.97s-3.636-3.77-4.97-2.485c-2.45 2.361 3.554 7.99 3.554 7.99"/>`),
		g.Group(children),
	)
}