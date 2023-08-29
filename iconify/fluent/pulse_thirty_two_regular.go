package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PulseThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.52 5a1 1 0 0 1 .944.735l4.578 16.646l3.5-11.668a1 1 0 0 1 1.881-.098L24.667 16H28a1 1 0 1 1 0 2h-4a1 1 0 0 1-.923-.615l-1.424-3.417l-3.695 12.32a1 1 0 0 1-1.922-.023L11.43 9.518l-2.477 7.785A1 1 0 0 1 8 18H4a1 1 0 1 1 0-2h3.269l3.278-10.303A1 1 0 0 1 11.52 5Z"/>`),
		g.Group(children),
	)
}