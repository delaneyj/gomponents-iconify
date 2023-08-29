package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSimpleFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 10.75A4.75 4.75 0 0 1 10.75 6h12v16.75H6v-12Zm0 14.5v12A4.75 4.75 0 0 0 10.75 42h12V25.25H6ZM25.25 42h12A4.75 4.75 0 0 0 42 37.25v-12H25.25V42ZM42 22.75v-12A4.75 4.75 0 0 0 37.25 6h-12v16.75H42Z"/>`),
		g.Group(children),
	)
}