package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SeekNext(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.402 12.5L5 6v13l1-.625l9.402-5.875Zm-1.887 0L6 17.196V7.804l7.515 4.696ZM18 6h-1v13h1V6Z"/>`),
		g.Group(children),
	)
}