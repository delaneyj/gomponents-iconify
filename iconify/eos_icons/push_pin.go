package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushPin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 13.04l-1.99-1.992V3.017h1v-2h-10v2h1v8.023h-.009L6 13.04v2h5.01L11 22h2l.01-6.96H18v-2z"/>`),
		g.Group(children),
	)
}