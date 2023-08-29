package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reorder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<circle cx="3.5" cy="3.5" r="1.5" fill="currentColor"/><circle cx="3.5" cy="8.5" r="1.5" fill="currentColor"/><circle cx="3.5" cy="13.5" r="1.5" fill="currentColor"/>`),
		g.Group(children),
	)
}