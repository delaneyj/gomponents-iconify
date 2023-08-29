package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inaturalist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.56 23.21c.88 0 3.06-8.39 12.94-5.05c-4.76 3.63-2.1 8.66-6.86 15.21A18.47 18.47 0 0 1 22 41.11a18.67 18.67 0 0 1-6.91-1.25A16 16 0 0 0 27 31.61a12.15 12.15 0 0 1-14-5.8a18.07 18.07 0 0 1 4.67-.69s-8.14-3.27-9-5.85c1.21-.95 4.94-.42 4.94-.42S5 14.1 4.5 10.86c15.25 0 25.23 12.35 26.06 12.35Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.32 18S19.15 11.42 19 6.89C22.91 7 30.29 13 31.27 22.41"/><circle cx="37.97" cy="19.87" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}