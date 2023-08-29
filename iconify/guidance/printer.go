package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Printer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M5.5 13.5c0 4.5-2 7.5-2 7.5v.5h17V21s-2-3-2-7.5m-13-6v-5h13v5M3 13.5h18m-18 4H.5v-10h23v10H21"/>`),
		g.Group(children),
	)
}