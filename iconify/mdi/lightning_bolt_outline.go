package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightningBoltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 9.47V11h3.76L13 14.53V13H9.24L11 9.47M13 1L6 15h5v8l7-14h-5V1Z"/>`),
		g.Group(children),
	)
}