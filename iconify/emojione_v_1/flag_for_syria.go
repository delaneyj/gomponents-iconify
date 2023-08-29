package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForSyria(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#e6e7e8" d="M0 25h64v14H0z"/><path fill="#ec1c24" d="M54 10H10C3.373 10 0 14.925 0 21v4h64v-4c0-6.075-3.373-11-10-11"/><path fill="#25333a" d="M0 43c0 6.075 3.373 11 10 11h44c6.627 0 10-4.925 10-11v-4H0v4"/><path fill="#137a08" d="m15.734 26.342l1.768 3.583l3.957.579l-2.862 2.796l.676 3.93l-3.539-1.86l-3.54 1.86l.678-3.93l-2.862-2.796l3.958-.579zm32.213 0l1.768 3.583l3.958.579L50.81 33.3l.68 3.93l-3.543-1.86l-3.537 1.86l.68-3.93l-2.87-2.796l3.96-.579z"/>`),
		g.Group(children),
	)
}