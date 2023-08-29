package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Expertsexchange(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.28.9H0L8.36 12L0 23.1h7.28L15.64 12zM24 .9h-7.28l-2.3 3.06l3.64 4.82zm-9.58 19.15l2.3 3.05H24l-5.94-7.88z"/>`),
		g.Group(children),
	)
}