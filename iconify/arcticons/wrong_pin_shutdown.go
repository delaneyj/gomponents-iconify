package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WrongPinShutdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="39" height="10.849" x="4.5" y="29.262" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.935"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.407 38.05v-6.727M21.154 9.57a7.812 7.812 0 1 0 5.692 0M24 7.889v7.865"/><circle cx="10.355" cy="34.687" r="1.682" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="34.687" r="1.682" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="17.177" cy="34.687" r="1.682" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}