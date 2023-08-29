package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AudioDocumentSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 10a1 1 0 1 0-2 0a1 1 0 0 0 2 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12Zm6.342 2.526A.5.5 0 0 1 7.9 4.2l.3.4A3.5 3.5 0 0 0 11 6v1a4.5 4.5 0 0 1-3-1.146V10a2 2 0 1 1-1-1.732V4.5a.5.5 0 0 1 .342-.474Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}