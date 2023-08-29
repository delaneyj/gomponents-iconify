package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ribbon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M371 8q19 0 33 13t13 34v676q0 12-11 17t-20-3L212 591q-4-3-7 0L31 745q-9 8-20 3T0 731V55q0-20 14-34T46 8h325z"/>`),
		g.Group(children),
	)
}