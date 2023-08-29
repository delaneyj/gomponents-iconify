package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Netflix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 2h4l2.94 8.83L13.5 2h4v20c-1.25-.22-2.63-.36-4.09-.42L10.5 13l-.07 8.59c-1.4.06-2.73.2-3.93.41V2Z"/>`),
		g.Group(children),
	)
}