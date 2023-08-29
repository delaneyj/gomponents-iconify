package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForCongoBrazzaville(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#c32129" d="M64 21v22c0 6.075-3.373 11-10 11H21.045l39.908-41.35c2.02 2.02 3.047 5 3.047 8.346"/><path fill="#28ae66" d="M0 43V21c0-6.075 3.373-11 10-11h33.272L3.204 51.509C1.076 49.491 0 46.43 0 43z"/><path fill="#f9cb38" d="M59.08 11.237C57.68 10.451 55.99 10 54 10H43.272L3.202 51.509a8.633 8.633 0 0 0 1.833 1.315C6.413 53.57 8.065 54 9.998 54h11.045l39.908-41.35a8.636 8.636 0 0 0-1.873-1.417"/>`),
		g.Group(children),
	)
}