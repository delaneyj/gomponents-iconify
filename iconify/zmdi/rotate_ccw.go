package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotateCcw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m138 126l139 138l-139 139L0 264zM60 264l78 78l78-78l-78-78zm334-133q57 56 57 135.5T394 402q-56 56-135 56q-49 0-93-24l32-31q29 13 61 13q62 0 105.5-44T408 266.5T364.5 161T259 117v69l-91-90l91-90v69q79 0 135 56z"/>`),
		g.Group(children),
	)
}