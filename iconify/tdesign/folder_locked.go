package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderLocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v5h-2V7H10.52l-2-2.5H3V19h10v2H1V2.5Zm18.5 11c.69 0 1.25.56 1.25 1.25v.75h-2.5v-.75c0-.69.56-1.25 1.25-1.25Zm3.25 2v-.75a3.25 3.25 0 0 0-6.5 0v.75h-1.251V22h9v-6.5H22.75Zm-.751 2V20h-5v-2.5h5Z"/>`),
		g.Group(children),
	)
}