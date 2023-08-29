package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArmouryCrate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 7.425L4.5 40.575l9.75-3.9L24 19.125l9.75 17.55l9.75 3.9Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 26.925l-5.85 9.75l5.85-1.95l5.85 1.95Z"/>`),
		g.Group(children),
	)
}