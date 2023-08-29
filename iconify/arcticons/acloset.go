package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Acloset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.382 27.638c10.876-3.617 23.766-.631 31.118-3.85"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 41.831c3.742.341 5.647-.533 7.454-4.645C20.99 18.895 29.8 4.328 36.854 6.276c-2.548 13.787-6.302 35.535.31 35.135l3.519-.213"/>`),
		g.Group(children),
	)
}