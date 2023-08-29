package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.29 8.71a1 1 0 0 0 1.42 0l4-4a1 1 0 1 0-1.42-1.42L17 6.59l-1.29-1.3a1 1 0 0 0-1.42 1.42ZM21 8a1 1 0 0 0-1 1v9a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8.41l5.88 5.89a3 3 0 0 0 2.11.87a3.08 3.08 0 0 0 2.16-.9l1.72-1.72a1 1 0 1 0-1.42-1.42l-1.75 1.75a1 1 0 0 1-1.4 0L5.41 7H11a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V9a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}