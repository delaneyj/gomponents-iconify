package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TescoGrocery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.306 41.011H9.694L5.5 19.297h37l-4.194 21.714zM32.074 6.989l-8.78 12.308M24 34.112v-7.917m6.99 7.917v-7.917m-14.071 7.917v-7.917"/>`),
		g.Group(children),
	)
}