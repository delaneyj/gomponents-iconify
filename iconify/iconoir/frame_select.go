package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrameSelect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5"><path d="M4.998 2H2v2.998h2.998V2Zm.001 1.501h14M3.5 4.999V19M20.498 5v14.002M4.999 20.501h14M4.998 19H2v2.998h2.998V19ZM21.997 2.002H19V5h2.998V2.002Zm0 17H19V22h2.998v-2.998Z"/><path d="m10.997 15.002l-3-7l7 3l-2.998.999l-1.002 3.001Z" clip-rule="evenodd"/><path d="m11.999 12.002l2.998 3l-2.998-3Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}