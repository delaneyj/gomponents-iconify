package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderImport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2.5h8.48l2 2.5H24v16H13v-2h9V7H11.52l-2-2.5H4v8.25H2V2.5Zm5.05 10.588l4.858 4.914l-4.858 4.914l-1.422-1.406l2.48-2.508H.997v-2h7.11l-2.48-2.508l1.422-1.406Z"/>`),
		g.Group(children),
	)
}