package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 4h1l3 3h7a2 2 0 0 1 2 2v8m-2 2H5a2 2 0 0 1-2-2V6a2 2 0 0 1 1.189-1.829M3 3l18 18"/>`),
		g.Group(children),
	)
}