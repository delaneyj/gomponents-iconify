package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.111 12a4.502 4.502 0 0 0 8.777 0H22v8a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-8h9.111ZM5 16h2v2H5v-2Zm10.5-2.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5ZM11.111 10H2V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v6h-2.111a4.502 4.502 0 0 0-8.777 0Z"/>`),
		g.Group(children),
	)
}