package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConsumptionO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M20.001 37.317c-9.549 0-17.317-7.769-17.317-17.317S10.452 2.683 20.001 2.683c9.548 0 17.315 7.769 17.315 17.317s-7.767 17.317-17.315 17.317zm0-33.134C11.279 4.183 4.184 11.278 4.184 20s7.096 15.817 15.817 15.817S35.816 28.721 35.816 20S28.722 4.183 20.001 4.183z"/><path fill="currentColor" d="m29.97 27.917l-2.916-1.827v1.081H12.631v-.684l4.586-3.611l3.309 2.888l5.301-6.658l.919.721l.933-4.298l-4.394 1.585l1.073.842l-4.075 5.121l-3.006-2.623l-4.646 3.66V13.143h1.081l-1.827-2.917l-1.825 2.917h1.079v15.521h15.915v1.081z"/>`),
		g.Group(children),
	)
}