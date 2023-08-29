package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceCreamOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.11 21.46L2.39 1.73L1.11 3l3.84 3.84C4.37 7.38 4 8.14 4 9c0 1.66 1.34 3 3 3l5 10l2.7-5.41l6.14 6.14l1.27-1.27M12 17.53l-3.11-6.22c.06-.05.11-.1.19-.15c.13.08.27.14.42.21l3.71 3.73L12 17.53M7.15 3.95C8.07 2.2 9.89 1 12 1c2.89 0 5.25 2.22 5.5 5.05C18.91 6.28 20 7.5 20 9c0 1.66-1.34 3-3 3l-.6 1.2l-9.25-9.25Z"/>`),
		g.Group(children),
	)
}