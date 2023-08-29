package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zig(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m23.53 1.02l-7.686 3.45h-7.06l-2.98 3.452h7.173L.47 22.98l7.681-3.607h7.065v-.002l2.978-3.45l-7.148-.001l12.482-14.9zM0 4.47v14.901h1.883l2.98-3.45H3.451v-8h.942l2.824-3.45H0zm22.117 0l-2.98 3.608h1.412v7.844h-.942l-2.98 3.45H24V4.47h-1.883z"/>`),
		g.Group(children),
	)
}