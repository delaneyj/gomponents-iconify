package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.5 12.04l-4.326 2.275L4 9.497L.5 6.086l4.837-.703L7.5 1l2.163 4.383l4.837.703L11 9.497l.826 4.818L7.5 12.041Z"/>`),
		g.Group(children),
	)
}