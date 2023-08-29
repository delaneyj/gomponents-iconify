package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZSystems(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 3h-8c-1.103 0-2 .897-2 2v22c0 1.103.897 2 2 2h8c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2zm0 6.92L20.85 5H27v4.92zM26.4 12L19 17.92V6.08L26.4 12zm.6 2.08l.001 11.84l-7.4-5.92l7.4-5.92zm-8 8L25.15 27H19v-4.92zM13 3H5c-1.103 0-2 .897-2 2v22c0 1.103.897 2 2 2h8c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2zm0 6.92L6.85 5H13v4.92zM12.4 12L5 17.92V6.08L12.4 12zm.6 2.08l.001 11.84l-7.4-5.92l7.4-5.92zm-8 8L11.15 27H5v-4.92z"/>`),
		g.Group(children),
	)
}