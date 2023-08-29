package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlightInternational(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 28a5.326 5.326 0 0 1-.671-.044A21.785 21.785 0 0 1 12.033 17H29.95c.024-.332.05-.663.05-1a14 14 0 1 0-14 14Zm11.95-13h-5.987A24.325 24.325 0 0 0 19.21 4.45A12.012 12.012 0 0 1 27.95 15ZM16.67 4.044A21.785 21.785 0 0 1 19.967 15h-7.934A21.785 21.785 0 0 1 15.33 4.044a5.159 5.159 0 0 1 1.342 0Zm-3.881.405A24.328 24.328 0 0 0 10.037 15H4.05a12.013 12.013 0 0 1 8.74-10.55ZM4.05 17h5.987a24.328 24.328 0 0 0 2.753 10.55A12.013 12.013 0 0 1 4.05 17Z"/><path fill="currentColor" fill-rule="evenodd" d="m25 25l5 2v-2l-5-2.5V20a1 1 0 0 0-2 0v2.5L18 25v2l5-2v3.5L21 30v1l3-1l3 1v-1l-2-1.5Z"/>`),
		g.Group(children),
	)
}