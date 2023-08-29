package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpThick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3.172L5.586 9.586a2 2 0 1 0 2.828 2.828L10 10.828v7.242a2 2 0 0 0 4 0v-7.242l1.586 1.586c.391.391.902.586 1.414.586s1.023-.195 1.414-.586a2 2 0 0 0 0-2.828L12 3.172z"/>`),
		g.Group(children),
	)
}