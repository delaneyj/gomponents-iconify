package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deviantart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1024 303L721 885l24 31h279v415H517l-44 30l-142 273l-30 30H0v-303l303-583l-24-30H0V333h507l44-30L693 30l30-30h301v303z"/>`),
		g.Group(children),
	)
}