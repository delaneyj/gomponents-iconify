package feather

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1l22 22m-6.28-11.94A10.94 10.94 0 0 1 19 12.55m-14 0a10.94 10.94 0 0 1 5.17-2.39m.54-5.11A16 16 0 0 1 22.58 9M1.42 9a15.91 15.91 0 0 1 4.7-2.88m2.41 9.99a6 6 0 0 1 6.95 0M12 20h.01"/>`),
		g.Group(children),
	)
}