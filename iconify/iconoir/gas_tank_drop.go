package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GasTankDrop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5" clip-rule="evenodd"><path stroke-width="1.493" d="M3 7.562A2.562 2.562 0 0 1 5.563 5H7V3h5v2h2.002A6.998 6.998 0 0 1 21 11.998v6.442a2.562 2.562 0 0 1-2.563 2.562H5.563A2.565 2.565 0 0 1 3 18.44V7.562Z"/><path d="M12 9s3 2.993 3 4.886c0 1.656-1.345 3-3 3c-1.656 0-2.988-1.344-3-3C9.01 11.992 12 9 12 9Z"/></g>`),
		g.Group(children),
	)
}