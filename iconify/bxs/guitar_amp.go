package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GuitarAmp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 6h-2V4c0-1.103-.897-2-2-2H8c-1.103 0-2 .897-2 2v2H4c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V8c0-1.103-.897-2-2-2zM8 4h8v2H8V4zM6 19a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm0-3a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm3 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm0-3a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm3 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm0-3a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm3 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm0-3a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm3 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm0-3a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm2-4H4V8h16v4z"/><path fill="currentColor" d="M14 9h2v2h-2zm3 0h2v2h-2z"/>`),
		g.Group(children),
	)
}