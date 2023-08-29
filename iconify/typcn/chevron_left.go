package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.414 5.586a2 2 0 0 0-2.828 0L5.171 12l6.415 6.414c.39.391.902.586 1.414.586s1.024-.195 1.414-.586a2 2 0 0 0 0-2.828L10.829 12l3.585-3.586a2 2 0 0 0 0-2.828z"/>`),
		g.Group(children),
	)
}