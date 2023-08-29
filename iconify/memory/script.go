package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Script(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M20 1H5v1H4v13h2V3h9v16h-2v-1h-1v-1H1v3h1v1h14v-1h1V3h2v2h2V2h-1"/>`),
		g.Group(children),
	)
}