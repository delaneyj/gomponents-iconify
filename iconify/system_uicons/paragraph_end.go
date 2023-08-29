package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphEnd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 5.5h13.978M3.5 7.5h13.978M3.5 9.5h13.978M3.5 11.5h13.978M3.5 13.5h13.978M3.5 15.5h7"/>`),
		g.Group(children),
	)
}