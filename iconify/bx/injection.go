package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Injection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.84 7.4L22.25 6L18 1.75l-1.4 1.41L18 4.58L16.6 6l-2.13-2.13l-1.41-1.42l-1.41 1.42l.7.71L3 13.92a2 2 0 0 0-.3 2.45L4 18.56l-2.25 2.28l1.41 1.41L5.44 20l2.19 1.31a2 2 0 0 0 1 .28a2 2 0 0 0 1.45-.59l9.34-9.34l.71.7l1.42-1.41l-1.42-1.42L18 7.4L19.42 6zM18 10.23l-9.34 9.35L6 18l-1.58-2.66l.86-.87l2.83 2.83l1.42-1.41l-2.83-2.83l1.41-1.41l2.83 2.82l1.41-1.41l-2.83-2.83l1.42-1.41l2.83 2.83l1.41-1.42l-2.83-2.83L13.77 6z"/>`),
		g.Group(children),
	)
}