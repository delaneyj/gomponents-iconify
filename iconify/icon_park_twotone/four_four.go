package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFourFour0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M25.745 44h-.896c-5.21 0-10.07-2.626-12.925-6.984l-2.195-3.35a5.629 5.629 0 0 1 1.072-7.382l2.288-1.936c.26-.22.411-.545.411-.887V7.25a3.25 3.25 0 0 1 6.5 0v-1a3.25 3.25 0 0 1 6.5 0v1a3.25 3.25 0 0 1 6.5 0v4a3.25 3.25 0 0 1 6.5 0v19.058c0 2.73-.838 5.417-2.38 7.671C34.556 41.726 30.285 44 25.745 44Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFourFour0)"/>`),
		g.Group(children),
	)
}