package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashAuto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 3h213l-85 192h85L64 451V259H0V3zm341 0l69 192h-41l-15-43h-68l-15 43h-41L299 3h42zm-46 120h50l-25-78z"/>`),
		g.Group(children),
	)
}