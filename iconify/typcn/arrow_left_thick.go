package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftThick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 11h-7.244l1.586-1.586a2 2 0 1 0-2.828-2.828L3.1 13l6.414 6.414c.39.391.902.586 1.414.586s1.023-.195 1.414-.586a2 2 0 0 0 0-2.828L10.756 15H18a2 2 0 0 0 0-4z"/>`),
		g.Group(children),
	)
}