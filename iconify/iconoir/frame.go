package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Frame(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5"><path d="M4.998 2.001H2v2.998h2.998V2.001Zm0 8.501H2V13.5h2.998v-2.998ZM20.498 5v5.503M3.5 5v5.503m16.998 2.999v5.502M3.5 13.502v5.502m1.499 1.498h5.5"/><path stroke-width="1.22" d="M4.999 3.503h5.5"/><path d="M13.498 20.499h5.5"/><path stroke-width="1.22" d="M13.498 3.501h5.5"/><path d="M4.998 19.001H2v2.998h2.998v-2.998ZM21.997 2.002H19V5h2.998V2.002ZM13.497 2H10.5v2.998h2.998V2Zm8.5 8.503H19V13.5h2.998v-2.998Zm0 8.499H19V22h2.998v-2.998Zm-8.5-.002H10.5v2.998h2.998V19Z"/></g>`),
		g.Group(children),
	)
}