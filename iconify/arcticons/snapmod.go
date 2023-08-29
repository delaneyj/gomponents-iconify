package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snapmod(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="21.7" height="39" x="13.15" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.022"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.644 14.009v0Z"/>`),
		g.Group(children),
	)
}