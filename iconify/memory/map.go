package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Map(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 4h2V3h2V2h4v1h3v1h2V3h2V2h3v16h-2v1h-2v1h-4v-1H9v-1H7v1H5v1H2V4m2 2v11h2v-1h1V5H5v1H4m8-1H9v11h1v1h3V6h-1V5m4 1h-1v11h2v-1h1V5h-2v1Z"/>`),
		g.Group(children),
	)
}