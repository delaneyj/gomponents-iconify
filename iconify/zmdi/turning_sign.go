package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurningSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M420 201q7 6 7 15t-7 15L228 423q-6 6-15 6t-15-6L6 231q-6-6-6-15t6-15L198 9q6-6 15-6t15 6zm-164 68l75-74l-75-75v53H149q-9 0-15 6.5t-6 15.5v85h43v-64h85v53z"/>`),
		g.Group(children),
	)
}