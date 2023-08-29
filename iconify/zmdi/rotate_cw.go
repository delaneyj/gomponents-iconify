package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotateCw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m312 126l139 138l-139 139l-138-139zm78 138l-78-78l-78 78l78 78zM56 131q56-56 136-56V6l90 90l-90 90v-69q-62 0-105.5 44T43 266.5T86.5 372T192 416q31 0 60-13l32 31q-43 24-92 24q-80 0-136-56T0 266.5T56 131z"/>`),
		g.Group(children),
	)
}