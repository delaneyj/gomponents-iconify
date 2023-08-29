package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatFlat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 171v42H149V85h192q36 0 61 25t25 61zM0 235h427v42H299v43H128v-43H0v-42zm109.5-41q-18.5 19-45 19.5T19 195T0 150t18.5-45t45-19.5t45.5 18t19 45t-18.5 45.5z"/>`),
		g.Group(children),
	)
}