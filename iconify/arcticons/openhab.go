package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openhab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.86 38.21a21.44 21.44 0 1 0-4.15-7.08"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.5 34.96l18.55-18.55l14.07 14.06"/>`),
		g.Group(children),
	)
}