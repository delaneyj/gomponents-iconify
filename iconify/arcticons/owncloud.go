package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Owncloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.09 13.73c5.84 0 7.44 4.48 7.44 6.8a10.37 10.37 0 0 1-.46 3.47c1.83-.45 6.46 0 6.46 4.73v.84a5.14 5.14 0 0 1-5.16 4.61H9.72a5.15 5.15 0 0 1-5.19-5v-.11a5.6 5.6 0 0 1 5.6-5.74c2.73 0 2.95.64 2.95.64s-.08-2.87 1.31-4.44a6.6 6.6 0 0 1 5.36-2.62a4.32 4.32 0 0 1 3.48 1.28s1-4.47 6.86-4.47Z"/>`),
		g.Group(children),
	)
}