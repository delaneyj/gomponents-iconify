package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M6 8.5h12M6 13h6M23.5 2H23c-3 .5-8 .75-11 .75S4 2.5 1 2H.5v21.5h.25l.154-.154A15.692 15.692 0 0 1 12 18.75c3 0 8 .25 11 .75h.5V2Z"/>`),
		g.Group(children),
	)
}