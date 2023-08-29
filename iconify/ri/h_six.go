package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.097 8l-2.598 4.5a4 4 0 1 1-3.453 1.981L18.788 8h2.309ZM4 4v7h7V4h2v16h-2v-7H4v7H2V4h2Zm14.5 10.5a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z"/>`),
		g.Group(children),
	)
}