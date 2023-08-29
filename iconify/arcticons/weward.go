package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Weward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.633 13.366C12.41 4.729-1.359 15.122 7.257 24.79l11.957 13.07l11.134-11.425c2.332-2.488 4.616-9.236-1.453-13.457c-3.204-2.003-7.336-1.634-10.262.388c-4.355 2.976-5.004 8.886-1.84 12.489L28.509 37.81l12.489-13.554c7.947-9.663-4.938-18.762-12.102-11.279"/><circle cx="24.055" cy="20.469" r="2.071" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}