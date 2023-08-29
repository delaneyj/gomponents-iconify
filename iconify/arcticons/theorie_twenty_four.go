package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TheorieTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.48 5.5h15.042v37H16.48z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.5 17.5H15l-7.999-8h9.5"/><circle cx="24" cy="13.501" r="4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="34.499" r="4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.5 28H15l-7.999-8h9.5M16.5 38.499H15l-7.999-8h9.5M31.5 17.5H33l7.999-8h-9.5M31.5 28H33l7.999-8h-9.5m.001 18.499H33l7.999-8h-9.5"/>`),
		g.Group(children),
	)
}