package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 8v5.81c-.88-.51-1.9-.81-3-.81c-3.31 0-6 2.69-6 6c0 .34.04.67.09 1H4a2 2 0 0 1-2-2V6c0-1.11.89-2 2-2h6l2 2h8a2 2 0 0 1 2 2m-2 8h-2v4h-2l3 3l3-3h-2v-4Z"/>`),
		g.Group(children),
	)
}