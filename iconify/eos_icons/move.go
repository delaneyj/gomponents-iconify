package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Move(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.017 12.008L18.019 8.01l-.002 2.993h-5.001v-4.99l2.993.002l-4.002-4.002l-3.998 3.998l2.994.001v4.991h-5L6.002 8.01l-3.998 3.998l4.002 4.002l-.002-2.993h4.999v4.993l-2.994.001l3.998 3.999l4.002-4.002l-2.993.001v-4.992h5l-.001 2.993l4.002-4.002z"/>`),
		g.Group(children),
	)
}