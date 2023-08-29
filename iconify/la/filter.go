package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 4v2.344l.219.281L13 16.344V28l1.594-1.188l4-3L19 23.5v-7.156l7.781-9.719l.219-.281V4zm2.281 2H24.72l-7.188 9H14.47zM15 17h2v5.5L15 24z"/>`),
		g.Group(children),
	)
}