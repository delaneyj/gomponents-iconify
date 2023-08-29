package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowImportTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21.25 4.5a.75.75 0 0 1 .743.648L22 5.25v13.5a.75.75 0 0 1-1.493.102l-.007-.102V5.25a.75.75 0 0 1 .75-.75Zm-9.04 1.887l.083-.094a1 1 0 0 1 1.32-.083l.094.083l4.997 4.998a1 1 0 0 1 .083 1.32l-.083.093l-4.996 5.004a1 1 0 0 1-1.499-1.32l.083-.094L15.581 13H3a1 1 0 0 1-.993-.883L2 12a1 1 0 0 1 .883-.993L3 11h12.584l-3.291-3.293a1 1 0 0 1-.083-1.32l.083-.094l-.083.094Z"/>`),
		g.Group(children),
	)
}