package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveFrame(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path stroke-miterlimit="1.5" d="M4.998 2H2v2.998h2.998V2Zm.001 1.5h14M3.5 4.998V19M20.498 5v14.002M4.999 20.5h14M4.998 19H2v2.998h2.998V19ZM21.997 2.001H19v2.998h2.998V2.001Zm0 17H19v2.998h2.998v-2.998Z"/><path d="M9 12h6"/></g>`),
		g.Group(children),
	)
}