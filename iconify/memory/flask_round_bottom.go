package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlaskRoundBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M11 11h2v2h-2v-2m2-10v2h1v5h1v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-1v1H8v-1H7v-1H6v-1H5v-1H4v-6h1v-1h1V9h1V8h1V3h1V1h4m-1 4h-2v4H9v1H8v1H7v1H6v2h1v-1h2v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h-1v-1h-1v-1h-1V9h-1V5Z"/>`),
		g.Group(children),
	)
}