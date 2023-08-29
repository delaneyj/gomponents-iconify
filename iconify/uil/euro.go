package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Euro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.2 17.41A6 6 0 0 1 14.46 20c-2.68 0-5-2-6-5H14a1 1 0 0 0 0-2H8.05c0-.33-.05-.67-.05-1s0-.67.05-1H14a1 1 0 0 0 0-2H8.47c1-3 3.31-5 6-5a6 6 0 0 1 4.73 2.59a1 1 0 1 0 1.6-1.18A7.92 7.92 0 0 0 14.46 2c-3.76 0-7 2.84-8.07 7H4a1 1 0 0 0 0 2h2.05v2H4a1 1 0 0 0 0 2h2.39c1.09 4.16 4.31 7 8.07 7a7.92 7.92 0 0 0 6.34-3.41a1 1 0 0 0-1.6-1.18Z"/>`),
		g.Group(children),
	)
}