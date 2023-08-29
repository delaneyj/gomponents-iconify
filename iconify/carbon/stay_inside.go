package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StayInside(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23 30h-2v-2a3.003 3.003 0 0 0-3-3h-4a3.003 3.003 0 0 0-3 3v2H9v-2a5.006 5.006 0 0 1 5-5h4a5.006 5.006 0 0 1 5 5zm-7-17a3 3 0 1 1-3 3a3 3 0 0 1 3-3m0-2a5 5 0 1 0 5 5a5 5 0 0 0-5-5z"/><path fill="currentColor" d="M30 30h-2V14.464L16 4.31L4 14.464V30H2V14a1 1 0 0 1 .354-.763l13-11a1 1 0 0 1 1.292 0l13 11A1 1 0 0 1 30 14Z"/>`),
		g.Group(children),
	)
}