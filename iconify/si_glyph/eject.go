package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.152 1.004a.989.989 0 0 0-.988.99v1.979c0 .547.442.99.988.99h13.855a.99.99 0 0 0 .99-.99V1.994a.99.99 0 0 0-.99-.99H2.152zm7.899 14.557a1.49 1.49 0 0 1-2.104 0L1.504 9.116c-.582-.58-.838-2.102 1-2.102h12.988c1.902 0 1.582 1.521 1.002 2.102l-6.443 6.445z"/>`),
		g.Group(children),
	)
}