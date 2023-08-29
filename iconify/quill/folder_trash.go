package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderTrash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M27 8c0 1.657-4.925 3-11 3S5 9.657 5 8m22 0c0-1.657-4.925-3-11-3S5 6.343 5 8m22 0l-3 18s-1 2-8 2s-8-2-8-2L5 8m13.5 8.5l-5 5m0-5l5 5"/>`),
		g.Group(children),
	)
}