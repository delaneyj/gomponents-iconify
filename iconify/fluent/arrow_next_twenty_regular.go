package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowNextTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.5 5a.5.5 0 0 1 .492.41L14 5.5v9a.5.5 0 0 1-.992.09L13 14.5v-9a.5.5 0 0 1 .5-.5Zm-7.854.146a.5.5 0 0 1 .638-.057l.07.057l4.5 4.5a.5.5 0 0 1 .057.638l-.057.07l-4.5 4.5a.5.5 0 0 1-.765-.638l.057-.07L9.793 10L5.646 5.854a.5.5 0 0 1 0-.708Z"/>`),
		g.Group(children),
	)
}