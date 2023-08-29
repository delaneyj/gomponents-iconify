package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDocumentSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12Zm3 2.497L9 4v1l-5-.003v-1Zm7 2.998H4v1h7v-1Zm0 3.006l-7-.008v1L11 11v-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}