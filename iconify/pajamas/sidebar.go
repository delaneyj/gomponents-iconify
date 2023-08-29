package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sidebar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 2.5h7a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-.5.5H6v-11Zm-1.5 0H3a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5h1.5v-11ZM1 3a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}