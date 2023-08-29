package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M3 18V6a2 2 0 0 1 2-2h4.539a2 2 0 0 1 1.562.75L12.2 6.126a1 1 0 0 0 .78.375H20a1 1 0 0 1 1 1V18a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1z"/><path stroke-linecap="round" d="M9.5 12.5h5M12 15v-5"/></g>`),
		g.Group(children),
	)
}