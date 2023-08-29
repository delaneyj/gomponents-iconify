package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Syringe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.71 2.29a1 1 0 0 0-1.42 0l-2.12 2.12l-.71-.7a1 1 0 0 0-1.41 0l-1.41 1.41l-.71-.71a1 1 0 0 0-1.42 0l-7.77 7.78l-.74-.7a1 1 0 0 0-1.38 1.41l3.53 3.54l-1.73 1.74l-.71-.72a1 1 0 0 0-1.42 1.42l2.83 2.83a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42l-.71-.7l1.74-1.74l3.53 3.53a1 1 0 0 0 .71.3a1 1 0 0 0 .7-1.71l-.7-.71l7.78-7.77a1 1 0 0 0 0-1.42l-.71-.71L20.29 8a1 1 0 0 0 0-1.41l-.7-.71l2.12-2.12a1 1 0 0 0 0-1.47ZM7.57 15l-1.42-1.39l1.41-1.42L9 13.61Zm2.82 2.83L9 16.44L10.39 15l1.42 1.42ZM13.22 15L9 10.78l4.24-4.24l.71.7l3.53 3.54Zm4.24-7l-1.41-1.46l.71-.71l.7.7l.7.7Z"/>`),
		g.Group(children),
	)
}