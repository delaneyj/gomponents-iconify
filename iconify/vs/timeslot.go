package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timeslot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M320 320q-120 0-220 103.5T0 667v324q0 140 100 246.5T320 1344h1024q119 0 219.5-103.5T1664 998V667q0-139-100.5-243T1344 320H320zm-64 320h1152v384H256V640z"/>`),
		g.Group(children),
	)
}