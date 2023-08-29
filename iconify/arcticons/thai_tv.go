package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThaiTv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.652 16.855h34.696c1.037 0 1.872.835 1.872 1.872l.28 9.916l-.28 9.916c-.03 1.036-.835 1.872-1.872 1.872H6.652a1.867 1.867 0 0 1-1.871-1.872S4.5 31.95 4.5 28.643s.28-9.916.28-9.916c0-1.037.835-1.871 1.872-1.871ZM4.986 37.576h38.028M4.986 33.18h38.028M4.986 19.71h38.028M4.986 24.106h38.028M22.19 16.594L14.305 7.57m11.505 9.024l7.886-9.025"/>`),
		g.Group(children),
	)
}