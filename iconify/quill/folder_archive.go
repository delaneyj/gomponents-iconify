package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderArchive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 11h20m-14 5h8M5 7v4h2v14.002C7 26.105 7.895 27 9 27h14c1.105 0 2-.894 2-1.998V11h2V7a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2Z"/>`),
		g.Group(children),
	)
}