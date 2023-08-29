package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpArrowCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0zm0 469q-88 0-150.5-62.5T43 256t62.5-150.5T256 43t150.5 62.5T469 256t-62.5 150.5T256 469zm13-315l-2-3h-22q-2 0-2 3l-64 42q-14 12-4 30t30 7l30-22v130q0 22 21 22t21-22V211l30 22q20 11 30-7q10-19-6-30z"/>`),
		g.Group(children),
	)
}