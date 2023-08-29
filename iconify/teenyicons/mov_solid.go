package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 7.5a.5.5 0 0 1 1 0v2a.5.5 0 0 1-1 0v-2Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM7.5 6A1.5 1.5 0 0 0 6 7.5v2a1.5 1.5 0 0 0 3 0v-2A1.5 1.5 0 0 0 7.5 6Zm-4.646.146A.5.5 0 0 0 2 6.5V11h1V7.707l.5.5l.5-.5V11h1V6.5a.5.5 0 0 0-.854-.354l-.646.647l-.646-.647ZM10 6h1v3.293l.5.5l.5-.5V6h1v3.707l-1.146 1.147a.5.5 0 0 1-.708 0L10 9.707V6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}