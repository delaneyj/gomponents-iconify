package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownwardsArrowToBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.152 11.004a.989.989 0 0 0-.988.99v1.979c0 .547.442.99.988.99h13.855a.99.99 0 0 0 .99-.99v-1.979a.99.99 0 0 0-.99-.99H1.152zm7.899-1.443a1.49 1.49 0 0 1-2.104 0L.504 3.116c-.582-.58-.838-2.102 1-2.102h12.988c1.902 0 1.582 1.521 1.002 2.102L9.051 9.561z"/>`),
		g.Group(children),
	)
}