package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="2"/><path d="M12 10.016A2.794 2.794 0 0 0 9.857 9C8.28 9 7 10.343 7 12s1.28 3 2.857 3c.854 0 1.62-.393 2.143-1.016M17 12c0 1.657-1.28 3-2.857 3c-1.578 0-2.857-1.343-2.857-3s1.279-3 2.857-3C15.72 9 17 10.343 17 12Z"/></g>`),
		g.Group(children),
	)
}