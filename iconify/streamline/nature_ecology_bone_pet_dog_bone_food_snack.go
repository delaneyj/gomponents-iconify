package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NatureEcologyBonePetDogBoneFoodSnack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 2.5a2 2 0 0 0-4 0a2 2 0 0 0 .59 1.41L3.91 8.09A2 2 0 0 0 2.5 7.5a2 2 0 0 0 0 4a2 2 0 0 0 4 0a2 2 0 0 0-.59-1.41l4.18-4.18a2 2 0 0 0 1.41.59a2 2 0 0 0 0-4Z"/>`),
		g.Group(children),
	)
}