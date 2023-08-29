package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTwoTwo0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M26.745 44h-.896c-5.21 0-10.07-2.626-12.925-6.984l-2.195-3.35a5.629 5.629 0 0 1 1.072-7.382l2.288-1.936c.26-.22.411-.545.411-.887V7.25a3.25 3.25 0 0 1 6.5 0v-1a3.25 3.25 0 0 1 6.5 0v11.5a3.25 3.25 0 0 1 6.5 0v4a3.25 3.25 0 0 1 6.5 0v10.219c0 1.649-.498 3.26-1.43 4.62l-.973 1.423A13.755 13.755 0 0 1 26.745 44Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTwoTwo0)"/>`),
		g.Group(children),
	)
}