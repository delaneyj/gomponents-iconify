package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WrenchCheckOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 1.09V6H7V1.09C4.16 1.57 2 4.03 2 7c0 2.22 1.21 4.15 3 5.19V21c0 .55.45 1 1 1h4c.55 0 1-.45 1-1v-8.81c1.79-1.04 3-2.97 3-5.19c0-2.97-2.16-5.43-5-5.91m1 9.37l-1 .58V20H7v-8.96l-1-.58C4.77 9.74 4 8.42 4 7c0-1 .37-1.94 1-2.65V8h6V4.35c.63.71 1 1.65 1 2.65c0 1.42-.77 2.74-2 3.46M21.6 13l1.4 1.41L16.47 21L13 17.5l1.4-1.41l2.07 2.08L21.6 13"/>`),
		g.Group(children),
	)
}