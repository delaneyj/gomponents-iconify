package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceDizzyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 2a14 14 0 1 0 14 14A14 14 0 0 0 16 2ZM9 16.41L7.59 15l2-2l-2-2L9 9.59l2 2l2-2L14.41 11l-2 2l2 2L13 16.41l-2-2ZM16 25a3 3 0 1 1 3-3a3 3 0 0 1-3 3Zm8.41-10L23 16.41l-2-2l-2 2L17.59 15l2-2l-2-2L19 9.59l2 2l2-2L24.41 11l-2 2Z"/>`),
		g.Group(children),
	)
}