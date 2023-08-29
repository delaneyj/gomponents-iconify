package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTwoWay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.8 10.167a.55.55 0 0 0-.779 0v.888H9.98c-.345-.073-2.035-.643-2.035-3.144c0-2.456 1.563-2.894 1.972-2.96H13v.903a.551.551 0 0 0 .779 0l2.06-1.427a.551.551 0 0 0 0-.779l-2.06-1.481a.55.55 0 0 0-.779 0v.847l-3.261.006c-1.198.131-3.314 1.176-3.711 4.022H1c-.553 0-1 .377-1 .929a1 1 0 0 0 1 1h5.044c.437 2.75 2.53 3.814 3.698 3.989l3.278.01v.863a.551.551 0 0 0 .779 0l2.039-1.405a.551.551 0 0 0 0-.779L13.8 10.167z"/>`),
		g.Group(children),
	)
}