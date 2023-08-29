package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Organization(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.01 10.99h-7v-2h-2v2H3.47v4h2v-2h5.54v2h2v-2h5.5v2h2v-4h-.5z"/><circle cx="12.01" cy="4.51" r="2.5" fill="currentColor"/><circle cx="4.47" cy="19.49" r="2.5" fill="currentColor"/><circle cx="12.01" cy="19.49" r="2.5" fill="currentColor"/><circle cx="19.51" cy="19.49" r="2.5" fill="currentColor"/>`),
		g.Group(children),
	)
}