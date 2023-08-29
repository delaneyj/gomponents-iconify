package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jblheadphones(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.525 30.123a18.497 18.497 0 1 1 34.949.004"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.231 40.467A8.333 8.333 0 0 1 8.434 27.772Zm9.538 0a8.333 8.333 0 1 0 10.797-12.695Zm-3.664-19.831a2.525 2.525 0 1 1 0 5.05H20.94v-10.1h4.166a2.525 2.525 0 1 1 0 5.05Zm0 0H20.94m-2.57-5.05v7.575a2.532 2.532 0 0 1-2.526 2.525h0a2.532 2.532 0 0 1-2.525-2.525v-.884m16.694-6.691v10.1h5.05"/>`),
		g.Group(children),
	)
}