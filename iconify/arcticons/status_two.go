package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatusTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.998 4.628a9.358 9.358 0 0 1 8.467 10.17c-.47 5.147-4.855 8.124-10.17 8.467c0 0-2.98.427-8.668-1.476a20.154 20.154 0 0 0-11.018-.31S17.153 2.785 34 4.628M14.002 43.372a9.358 9.358 0 0 1-8.467-10.17c.47-5.147 4.855-8.124 10.17-8.467c0 0 2.98-.427 8.668 1.476a20.154 20.154 0 0 0 11.018.31s-4.544 18.694-21.39 16.851"/>`),
		g.Group(children),
	)
}