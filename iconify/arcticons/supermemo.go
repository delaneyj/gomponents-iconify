package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Supermemo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32 41.56c-2.11-3.6-2.38-5.15-2.38-6.82s5.26-6.43 5.26-11a10.94 10.94 0 0 0-11.1-10.48c-7.13 0-10 5.16-10 8.93s-1.94 5-1.94 5.82s1.71 1.44 1.71 1.44a10 10 0 0 0 .23 2.88c.31.92-.95 3.55 2.66 3.55c1 0 1-.56 2.21-.56s2.17 2.11 2.17 4A14.22 14.22 0 0 1 20 43.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.14 22.25c2.66 2.81 7 7.5 6.58 15.63m2.85 2.18c-1.26-2.69-2.11-9.12 1.14-16.44m-8.31-1.98c.92 0 .92.79 1.85.79s.93-.79 1.86-.79s.93.79 1.85.79s.93-.79 1.86-.79M24 10.04V4.5m7.48 7.42l3.91-3.91m1.32 9.44h5.54m-25.73-5.53l-3.91-3.91m-1.32 9.44H5.75"/>`),
		g.Group(children),
	)
}