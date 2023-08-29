package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotepadDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M208 48v152a24 24 0 0 1-24 24H72a24 24 0 0 1-24-24V48a8 8 0 0 1 8-8h144a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M168 128a8 8 0 0 1-8 8H96a8 8 0 0 1 0-16h64a8 8 0 0 1 8 8Zm-8 24H96a8 8 0 0 0 0 16h64a8 8 0 0 0 0-16Zm56-104v152a32 32 0 0 1-32 32H72a32 32 0 0 1-32-32V48a16 16 0 0 1 16-16h16v-8a8 8 0 0 1 16 0v8h32v-8a8 8 0 0 1 16 0v8h32v-8a8 8 0 0 1 16 0v8h16a16 16 0 0 1 16 16Zm-16 0h-16v8a8 8 0 0 1-16 0v-8h-32v8a8 8 0 0 1-16 0v-8H88v8a8 8 0 0 1-16 0v-8H56v152a16 16 0 0 0 16 16h112a16 16 0 0 0 16-16Z"/></g>`),
		g.Group(children),
	)
}