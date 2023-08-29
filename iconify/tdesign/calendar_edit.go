package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 1v3h8V1h2v3h4v2.001h.001l-.001 6L20 12v-1H4v9h8v2H2V4h4V1h2ZM4 9h16V6H4v3Zm15.287 4.086l4.127 4.127l-5.286 5.287H14v-4.128l5.287-5.286Zm-.672 3.5l1.299 1.3l.672-.673l-1.3-1.299l-.671.672Zm-.115 2.713L17.2 18L16 19.2v1.3h1.3l1.2-1.2Z"/>`),
		g.Group(children),
	)
}