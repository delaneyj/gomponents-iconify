package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M6 2h2V1h6v1h2v1h1v1h1v1h1v2h1v7h-1v2h-1v4h-1v1H5v-1H4v-4H3v-2H2V8h1V5h1V4h1V3h1V2m9 3V4h-2V3H9v1H7v1H6v1H5v3H4v4h1v2h1v4h2v-2h2v2h2v-2h2v2h2v-4h1v-2h1V8h-1V6h-1V5h-1M7 8h3v3H7V8m5 3V8h3v3h-3m-2 2h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}