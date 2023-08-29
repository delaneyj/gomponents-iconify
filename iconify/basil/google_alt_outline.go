package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 12c0-5.384 4.366-9.75 9.75-9.75a9.71 9.71 0 0 1 6.64 2.623a.75.75 0 0 1 .018 1.08l-2.545 2.545a.75.75 0 0 1-1.029.03A4.621 4.621 0 0 0 12 7.35a4.65 4.65 0 0 0 0 9.3a4.637 4.637 0 0 0 3.883-2.1H12a.75.75 0 0 1-.75-.75v-3.6a.75.75 0 0 1 .75-.75h8.825a.75.75 0 0 1 .736.604A10.2 10.2 0 0 1 21.75 12c0 5.384-4.366 9.75-9.75 9.75S2.25 17.384 2.25 12ZM12 3.75a8.25 8.25 0 1 0 8.185 7.2H12.75v2.1h4.336a.75.75 0 0 1 .707 1A6.148 6.148 0 0 1 5.85 12A6.15 6.15 0 0 1 12 5.85c1.313 0 2.527.415 3.524 1.116l1.499-1.5A8.187 8.187 0 0 0 12 3.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}