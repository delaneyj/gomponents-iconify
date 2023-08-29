package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1043 301L512 832l531 531q19 19 19 45t-19 45l-166 166q-19 19-45 19t-45-19L45 877q-19-19-19-45t19-45L787 45q19-19 45-19t45 19l166 166q19 19 19 45t-19 45z"/>`),
		g.Group(children),
	)
}