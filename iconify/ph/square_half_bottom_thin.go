package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareHalfBottomThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 44H56a12 12 0 0 0-12 12v144a12 12 0 0 0 12 12h144a12 12 0 0 0 12-12V56a12 12 0 0 0-12-12ZM56 52h144a4 4 0 0 1 4 4v68H52V56a4 4 0 0 1 4-4Zm52 80v72H84v-72Zm8 0h24v72h-24Zm32 0h24v72h-24Zm-96 68v-68h24v72H56a4 4 0 0 1-4-4Zm148 4h-20v-72h24v68a4 4 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}