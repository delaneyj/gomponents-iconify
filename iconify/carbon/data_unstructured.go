package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataUnstructured(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 24a2 2 0 1 1-2 2a2 2 0 0 1 2-2m0-2a4 4 0 1 0 4 4a4 4 0 0 0-4-4zM16 4a2 2 0 1 1-2 2a2 2 0 0 1 2-2m0-2a4 4 0 1 0 4 4a4 4 0 0 0-4-4zm10 2a2 2 0 1 1-2 2a2 2 0 0 1 2-2m0-2a4 4 0 1 0 4 4a4 4 0 0 0-4-4zm-8 22v4h-4v-4h4m2-2h-8v8h8z"/><path fill="currentColor" d="M27 22.14V17a2 2 0 0 0-2-2H7v-5h3V2H2v8h3v5a2 2 0 0 0 2 2h18v5.14a4 4 0 1 0 2 0ZM4 4h4v4H4Zm22 24a2 2 0 1 1 2-2a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}