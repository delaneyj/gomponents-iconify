package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 3.5h6a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2zm8 2.5h1a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-1zm-10 0h-1a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h1z"/>`),
		g.Group(children),
	)
}