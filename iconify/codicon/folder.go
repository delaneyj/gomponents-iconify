package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.5 3H7.71l-.85-.85L6.51 2h-5l-.5.5v11l.5.5h13l.5-.5v-10L14.5 3zm-.51 8.49V13h-12V7h4.49l.35-.15l.86-.86H14v1.5l-.01 4zm0-6.49h-6.5l-.35.15l-.86.86H2v-3h4.29l.85.85l.36.15H14l-.01.99z"/>`),
		g.Group(children),
	)
}