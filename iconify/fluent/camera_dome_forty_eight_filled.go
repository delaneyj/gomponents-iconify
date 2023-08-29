package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraDomeFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 7.25A3.25 3.25 0 0 1 7.25 4h33.5a3.25 3.25 0 0 1 0 6.5H7.25A3.25 3.25 0 0 1 4 7.25ZM24 18.5a8.5 8.5 0 1 0 0 17a8.5 8.5 0 0 0 0-17Zm0 2.5a6 6 0 1 1 0 12a6 6 0 0 1 0-12Zm18-8H6v13c0 9.941 8.059 18 18 18s18-8.059 18-18V13ZM13 27c0-6.075 4.925-11 11-11s11 4.925 11 11s-4.925 11-11 11s-11-4.925-11-11Z"/>`),
		g.Group(children),
	)
}