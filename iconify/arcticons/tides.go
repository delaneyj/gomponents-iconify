package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tides(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.313 34.635C1.201 24 9.16 13.812 18.123 12.265c18.578-3.204 19.318 13.679 19.318 13.679"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.443 25.944c-4.1-11.683-16.221-9.095-21.518-3.798C10.81 27.261 10.147 41.243 24 45.5"/>`),
		g.Group(children),
	)
}