package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowPreviousTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 5a.5.5 0 0 0-.492.41L5.5 5.5v9a.5.5 0 0 0 .992.09l.008-.09v-9A.5.5 0 0 0 6 5Zm7.854.146a.5.5 0 0 0-.638-.057l-.07.057l-4.5 4.5a.5.5 0 0 0-.057.638l.057.07l4.5 4.5a.5.5 0 0 0 .765-.638l-.057-.07L9.707 10l4.147-4.146a.5.5 0 0 0 0-.708Z"/>`),
		g.Group(children),
	)
}