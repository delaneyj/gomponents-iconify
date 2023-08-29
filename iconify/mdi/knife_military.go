package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KnifeMilitary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 2l-4.61 1.75l-6.93 6.93L14 14.22l6.92-6.93C22.43 5.78 22 2 22 2M8.33 10l-1.41 1.39l1.41 1.41l-5.65 5.66L6.21 22l5.66-5.66l1.41 1.42l1.42-1.42L8.33 10Z"/>`),
		g.Group(children),
	)
}