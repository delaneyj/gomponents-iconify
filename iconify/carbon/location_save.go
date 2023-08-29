package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationSave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18 30h2v-2a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v2h2v-2a3.003 3.003 0 0 0-3-3h-6a3.003 3.003 0 0 0-3 3zm6-6a4 4 0 1 1 4-4a4.004 4.004 0 0 1-4 4zm0-6a2 2 0 1 0 2 2a2.002 2.002 0 0 0-2-2zm1.342-7.94l-22-8A1 1 0 0 0 2.06 3.341l8 22a1.001 1.001 0 0 0 .924.659H11a1 1 0 0 0 .929-.629l3.841-9.602l9.602-3.84a1 1 0 0 0-.03-1.87zm-10.714 4.01l-.398.16l-.159.398l-3.027 7.568L4.672 4.672l17.524 6.372z"/>`),
		g.Group(children),
	)
}