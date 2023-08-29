package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForUnitedArabEmirates(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#e6e7e8" d="M14 26h50v13H14z"/><path fill="#25333a" d="M14 54h40c6.627 0 10-4.925 10-11v-4H14v15"/><path fill="#137a08" d="M54 10H14v16h50v-5c0-6.075-3.373-11-10-11"/><path fill="#ec1c24" d="M14 39V10h-4C3.373 10 0 14.925 0 21v22c0 6.075 3.373 11 10 11h4V39z"/>`),
		g.Group(children),
	)
}