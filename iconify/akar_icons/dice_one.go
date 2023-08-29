package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiceOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" stroke-linejoin="round" rx="2"/><path d="M12.25 11.75v.5M12 12h.5"/></g>`),
		g.Group(children),
	)
}