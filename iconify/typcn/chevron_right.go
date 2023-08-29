package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.586 5.586a2 2 0 0 0 0 2.828L12.171 12l-3.585 3.586a2 2 0 1 0 2.828 2.828L17.829 12l-6.415-6.414a2 2 0 0 0-2.828 0z"/>`),
		g.Group(children),
	)
}