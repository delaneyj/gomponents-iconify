package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShiftmanagerPro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.256 13.015L20.763 39.851H14.75L5.5 23.83l9.25-16.022h4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.987 23.83l9.25-16.022h6.013L42.5 23.83l-9.25 16.021h-4"/>`),
		g.Group(children),
	)
}