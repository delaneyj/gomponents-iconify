package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bkash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.981 8.632s-4.163 14.704-3.809 14.704s16.476 2.923 16.476 2.923Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.981 8.632L6.329 6.152l12.843 17.184l2.215 10.186l14.261-7.263l3.72-8.814l-8.975 1.501m7.536 1.909H43l-3.632-3.41"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.387 33.522l.354 1.905l-8.016 6.421l5.447-18.512m16.476 2.923l-.531 2.879l-12.269 3.64M8.455 8.997H5L16.044 19.15"/>`),
		g.Group(children),
	)
}