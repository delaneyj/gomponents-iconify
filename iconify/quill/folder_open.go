package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M4 26V8a2 2 0 0 1 2-2h6c3 0 3 3 5 3h7a2 2 0 0 1 2 2v2M4 26l3.783-12.294A1 1 0 0 1 8.739 13H26M4 26h19.523a2 2 0 0 0 1.911-1.412l3.168-10.294A1 1 0 0 0 27.646 13H26"/>`),
		g.Group(children),
	)
}