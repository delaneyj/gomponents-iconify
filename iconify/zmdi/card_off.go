package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m354 43l1 249L113 50l50-50h149q17 0 29.5 12.5T354 43zM27 19l373 372l-27 28l-40-41q-10 6-21 6H99q-18 0-30.5-12.5T56 341V102L0 46z"/>`),
		g.Group(children),
	)
}