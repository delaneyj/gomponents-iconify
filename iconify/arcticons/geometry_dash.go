package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GeometryDash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m3.471 23.994l20.5-20.5l20.498 20.5L23.97 44.493z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.07 30.91l13.853-13.852l3.323 3.324l-13.852 13.852zm-4.73-8.466l3.429-3.43l3.429 3.43l-3.43 3.429zm6.687-6.684l3.43-3.43l3.429 3.43l-3.43 3.43z"/>`),
		g.Group(children),
	)
}