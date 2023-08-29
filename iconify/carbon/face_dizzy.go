package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceDizzy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 2a14 14 0 1 0 14 14A14 14 0 0 0 16 2Zm0 26a12 12 0 1 1 12-12a12 12 0 0 1-12 12Z"/><path fill="currentColor" d="M24.41 11L23 9.59l-2 2l-2-2L17.59 11l2 2l-2 2L19 16.41l2-2l2 2L24.41 15l-2-2l2-2zm-10 0L13 9.59l-2 2l-2-2L7.59 11l2 2l-2 2L9 16.41l2-2l2 2L14.41 15l-2-2l2-2zM16 19a3 3 0 1 0 3 3a3 3 0 0 0-3-3z"/>`),
		g.Group(children),
	)
}