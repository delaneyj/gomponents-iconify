package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diigo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M134 177q27 0 50.5-10.5t36-20.5t31.5-29v345h-57q-1-14-1-42v-97q0-18 .5-54t.5-54q-4 1-14 4.5t-14 4.5q-55 14-95-10v248H14V117q27 29 54 45t66 15zm91-100q15 0 26-11t11-27t-11-27t-26-11q-16 0-26.5 11T188 39t10.5 27T225 77zM43 77q16 0 26.5-11T80 39T69.5 12T43 1Q28 1 17 12T6 39t11 27t26 11z"/>`),
		g.Group(children),
	)
}