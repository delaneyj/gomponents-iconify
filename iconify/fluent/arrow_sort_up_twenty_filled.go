package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSortUpTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.75 3c.235 0 .452.108.59.278l2.94 2.945a.75.75 0 0 1 0 1.061a.748.748 0 0 1-1.058 0L10.5 5.566V16.25a.75.75 0 0 1-1.5 0V5.556L7.278 7.289a.748.748 0 0 1-1.059 0a.75.75 0 0 1 0-1.061L9.223 3.22A.731.731 0 0 1 9.75 3Z"/>`),
		g.Group(children),
	)
}