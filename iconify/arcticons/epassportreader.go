package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Epassportreader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="39" height="25" x="4.5" y="11.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.999"/><circle cx="24" cy="24" r="6.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.5 24h-13m26 0h13"/>`),
		g.Group(children),
	)
}