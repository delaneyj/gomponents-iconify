package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Teamviewer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.101 24l13.737 6.988l-1.747-4.728H24M5.101 24l13.737-6.988l-1.747 4.728H24m0 4.52h6.909l-1.747 4.728L42.899 24m0 0l-13.737-6.988l1.747 4.728H24"/>`),
		g.Group(children),
	)
}