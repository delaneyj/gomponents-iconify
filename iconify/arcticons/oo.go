package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="25.5" cy="21.3" r="2.9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.35 2.846a19.713 19.713 0 0 0 24.45 26.47"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.653 45.104a19.736 19.736 0 0 0-18.2-27.52l.007.016a19.802 19.802 0 0 0-6.25 1.05"/>`),
		g.Group(children),
	)
}