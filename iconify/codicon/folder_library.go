package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderLibrary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M7.71 3h6.79l.51.5V7H14V5.99H7.69l-.86.86l-.35.15H1.99v6H7v1H1.51l-.5-.5v-11l.5-.5h5l.35.15l.85.85Zm-.22 2h6.5l.01-.99H7.5l-.36-.15l-.85-.85H2v3h4.28l.86-.86l.35-.15Z" clip-rule="evenodd"/><path d="M8 8h1v6H8zm2 0h1v6h-1zm2.004.352l.94-.342l2.052 5.638l-.94.342z"/></g>`),
		g.Group(children),
	)
}