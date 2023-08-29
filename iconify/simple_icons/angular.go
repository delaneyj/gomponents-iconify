package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Angular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.96 12.648h4.08L12 7.74l-2.04 4.908zM12 0L.828 3.984l1.704 14.772L12 24l9.468-5.244l1.704-14.772L12 0zm6.972 18.312h-2.604l-1.404-3.504H9.036l-1.404 3.504H5.028L12 2.652l6.972 15.66z"/>`),
		g.Group(children),
	)
}