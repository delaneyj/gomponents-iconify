package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WireFlow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 14.117l12.962 19.766m12.962-19.766l-4.041 6.163m-4.992 7.612l-3.93 5.991m.115-19.766l12.962 19.766M43.5 14.117L30.538 33.883"/>`),
		g.Group(children),
	)
}