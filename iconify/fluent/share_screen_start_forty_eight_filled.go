package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareScreenStartFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.25 8A4.25 4.25 0 0 0 4 12.25v23.5A4.25 4.25 0 0 0 8.25 40h31.5A4.25 4.25 0 0 0 44 35.75v-23.5A4.25 4.25 0 0 0 39.75 8H8.25ZM24 14.75c.331 0 .65.132.884.366l6 6a1.25 1.25 0 0 1-1.768 1.768l-3.866-3.866V31.75a1.25 1.25 0 1 1-2.5 0V19.018l-3.866 3.866a1.25 1.25 0 0 1-1.768-1.768l6-6A1.25 1.25 0 0 1 24 14.75Z"/>`),
		g.Group(children),
	)
}