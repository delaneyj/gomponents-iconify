package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurnOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M11.124 1.613v2.349c1.606.887 2.632 2.624 2.632 4.578c0 2.854-2.35 5.169-5.248 5.169c-2.896 0-5.246-2.314-5.246-5.169c0-2.017 1.012-3.749 2.696-4.601V1.611c-2.937.975-4.895 3.693-4.895 6.929c0 4.052 3.334 7.335 7.444 7.335c4.111 0 7.446-3.283 7.446-7.335c.001-3.177-1.973-5.902-4.829-6.927z"/><path d="M8.46 7.873c.778 0 1.412-.48 1.412-1.075V1.115C9.872.521 9.238.04 8.46.04c-.779 0-1.412.481-1.412 1.075v5.683c0 .595.633 1.075 1.412 1.075z"/></g>`),
		g.Group(children),
	)
}