package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jeevansathi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.64 11l-5.832 21.77a5.708 5.708 0 0 1-11.028-2.955l.517-1.93m16.849 4.737a5.206 5.206 0 0 0 3.524.88h.88a2.863 2.863 0 1 0 0-5.725h-1.98a2.863 2.863 0 1 1 0-5.726h.881c1.982 0 2.863.22 3.524.88"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.6v20.9a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v12.1"/>`),
		g.Group(children),
	)
}