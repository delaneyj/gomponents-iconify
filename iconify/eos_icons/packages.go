package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Packages(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.04 6.01h-16v2h16v-2zm-2.02-4H6.01v2h12.01v-2zM2 10v2h2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8h2v-2Zm9 9H6v-2h5Z"/>`),
		g.Group(children),
	)
}