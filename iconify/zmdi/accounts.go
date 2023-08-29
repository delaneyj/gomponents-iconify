package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Accounts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M320 171q-27 0-45.5-19T256 106.5t18.5-45T320 43t45.5 18.5t18.5 45t-18.5 45.5t-45.5 19zm-170.5 0q-26.5 0-45.5-19t-19-45.5t19-45T149.5 43t45 18.5t18.5 45t-18.5 45.5t-45 19zm0 42q27.5 0 60.5 8t61 26t28 41v53H0v-53q0-23 27.5-41t61-26t61-8zm170.5 0q28 0 61 8t60.5 26t27.5 41v53H341v-53q0-43-42-74q13-1 21-1z"/>`),
		g.Group(children),
	)
}