package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendingUpRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.7 17.625q-.3-.3-.288-.713t.288-.687l5.275-5.35Q8.55 10.3 9.4 10.3t1.425.575l2.575 2.6l5.2-5.15H17q-.425 0-.713-.288T16 7.325q0-.425.288-.713T17 6.325h4q.425 0 .713.288t.287.712v4q0 .425-.288.713t-.712.287q-.425 0-.713-.287T20 11.325v-1.6L14.825 14.9q-.575.575-1.425.575t-1.425-.575L9.4 12.325l-5.3 5.3q-.275.275-.7.275t-.7-.275Z"/>`),
		g.Group(children),
	)
}