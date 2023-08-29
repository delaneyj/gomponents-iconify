package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeVariantOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5.64 3.64l15.72 15.72l-1.41 1.42L16 16.83V20l-5-5H7V9h1.17L4.22 5.05l1.42-1.41M16 4v7.17l-3.59-3.59L16 4Z"/>`),
		g.Group(children),
	)
}