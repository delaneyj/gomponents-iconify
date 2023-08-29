package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bluetooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="m8 11.967l.969-.943L8 10v1.967zm.03-5.985l.939-.945l-.939-1.028v1.973z"/><path d="M7.506 0C5.033 0 3.027 1.546 3.027 3.453v9.095C3.027 14.454 5.033 16 7.506 16c2.472 0 4.477-1.546 4.477-3.452V3.453C11.982 1.546 9.978 0 7.506 0zm2.946 11.011L6.94 14V8.815l-2.22 2.066l-.674-.748l2.09-2.07L4 6.1l.72-.666l2.22 1.912V2l3.56 3.117l-2.862 2.978l2.814 2.916z"/></g>`),
		g.Group(children),
	)
}