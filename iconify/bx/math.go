package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Math(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 2H5v3H2v2h3v3h2V7h3V5H7V2zm7 3h8v2h-8zm0 10h8v2h-8zm0 4h8v2h-8zm-5.71-4.71L6 16.59l-2.29-2.3l-1.42 1.42L4.59 18l-2.3 2.29l1.42 1.42L6 19.41l2.29 2.3l1.42-1.42L7.41 18l2.3-2.29l-1.42-1.42z"/>`),
		g.Group(children),
	)
}