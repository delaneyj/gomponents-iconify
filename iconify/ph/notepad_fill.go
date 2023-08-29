package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotepadFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 32h-16v-8a8 8 0 0 0-16 0v8h-32v-8a8 8 0 0 0-16 0v8H88v-8a8 8 0 0 0-16 0v8H56a16 16 0 0 0-16 16v152a32 32 0 0 0 32 32h112a32 32 0 0 0 32-32V48a16 16 0 0 0-16-16Zm-40 136H96a8 8 0 0 1 0-16h64a8 8 0 0 1 0 16Zm0-32H96a8 8 0 0 1 0-16h64a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}