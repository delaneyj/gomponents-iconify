package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PocketCastsAltThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 45.5C12.126 45.5 2.5 35.874 2.5 24S12.126 2.5 24 2.5S45.5 12.126 45.5 24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 37c-7.18 0-13-5.82-13-13s5.82-13 13-13s13 5.82 13 13m.5 7.5v12m6-6h-12"/>`),
		g.Group(children),
	)
}