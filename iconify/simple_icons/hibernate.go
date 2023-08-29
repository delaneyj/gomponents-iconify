package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hibernate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.365 0L9.98 7.994h8.95L14.31 0H5.366zm-.431.248L.46 7.994l4.613 8.008L9.55 8.24L4.934.248zm13.992 7.75l-4.475 7.76l4.617 7.992l4.471-7.744l-4.613-8.008zm-4.905 8.006l-8.95.002L9.688 24h8.946l-4.615-7.994l.001-.002Z"/>`),
		g.Group(children),
	)
}