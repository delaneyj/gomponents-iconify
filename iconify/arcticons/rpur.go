package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rpur(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.69 28.728l-5.006-.01l-2.494-4.34l2.511-4.33l5.005.01l2.494 4.34Zm8.35 13.717l5.46-9.247l-3.663-2.232l-4.029 7.118l-8.334.07l.045 4.29h10.52m.008-36.84l-10.524-.049l-.05 4.29l7.964-.096l4.315 7.13l3.665-2.228l-5.37-9.047M5.5 23.82l5.003 8.788l3.75-2.08l-3.86-6.492l3.86-6.923l-3.75-2.081L5.5 23.82"/>`),
		g.Group(children),
	)
}