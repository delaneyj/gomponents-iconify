package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GreaterThan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.636 4.372L16.985 12l-13.35 7.628l.993 1.736L21.016 12L4.628 2.636l-.992 1.736Z"/>`),
		g.Group(children),
	)
}