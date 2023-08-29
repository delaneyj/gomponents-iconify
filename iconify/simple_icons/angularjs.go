package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Angularjs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.964 0L.672 3.974l1.784 14.794L11.976 24l9.568-5.303l1.784-14.794zm-.027 1.258l10.265 3.5l-1.663 13.232l-8.602 4.76l-8.469-4.697L1.939 4.822zm0 .78L4.957 17.57l2.604-.048l1.4-3.501h6.257l1.532 3.55l2.492.046zm.02 4.98l2.355 4.93H9.878Z"/>`),
		g.Group(children),
	)
}