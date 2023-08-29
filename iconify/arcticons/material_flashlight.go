package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaterialFlashlight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="19.033" cy="29.294" r="3.506" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.033 9.427l18.699 18.698m-18.699-7.947L6.178 33.035L15.41 42.5l12.856-14.024l9.7-.304l3.856-3.973L23.123 5.5l-4.207 4.09Z"/>`),
		g.Group(children),
	)
}