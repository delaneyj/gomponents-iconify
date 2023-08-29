package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FdroidNearby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="22.671" cy="23.916" r="8.351" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="39.136" cy="16.656" r="4.364" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="9.955" cy="36.631" r="5.455" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.177 41.58a18.022 18.022 0 0 0 21.38-15.6M33.686 9.676a18.011 18.011 0 0 0-28.66 17.827"/>`),
		g.Group(children),
	)
}