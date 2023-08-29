package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlugX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.55 17.733a5.806 5.806 0 0 1-7.356-4.052a5.81 5.81 0 0 1 1.537-5.627L9.785 6l7.165 7.165M4 20l3.5-3.5M15 4l-3.5 3.5M20 9l-3.5 3.5M16 16l4 4m0-4l-4 4"/>`),
		g.Group(children),
	)
}