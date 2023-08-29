package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tidal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.012 3.992L8.008 7.996L4.004 3.992L0 7.996L4.004 12l4.004-4.004L12.012 12l-4.004 4.004l4.004 4.004l4.004-4.004L12.012 12l4.004-4.004l-4.004-4.004zm4.03 4.004l3.979-3.979L24 7.996l-3.979 3.979z"/>`),
		g.Group(children),
	)
}