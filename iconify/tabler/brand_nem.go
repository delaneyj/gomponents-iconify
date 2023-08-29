package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandNem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12.182 2c1.94.022 3.879.382 5.818 1.08l.364.135A23.075 23.075 0 0 1 22 5c0 5.618-1.957 10.258-5.87 13.92c-1.24 1.239-2.5 2.204-3.78 2.898L12 22c-1.4-.703-2.777-1.729-4.13-3.079C3.958 15.258 2 10.618 2 5c2.545-1.527 5.09-2.471 7.636-2.832L10 2.12A16.786 16.786 0 0 1 11.818 2h.364z"/><path d="M2.1 7.07C4.173 13.79 7.473 14.767 12 10c0-4 1.357-6.353 4.07-7.06l.59-.11m-.31 15.68S19 13 12 10"/></g>`),
		g.Group(children),
	)
}