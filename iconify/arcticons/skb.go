package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="8.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5C12.126 2.5 2.5 12.126 2.5 24S12.126 45.5 24 45.5S45.5 35.874 45.5 24"/><circle cx="36" cy="12" r="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}