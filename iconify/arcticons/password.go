package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Password(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 25.28a3.26 3.26 0 0 0-1.64 6.07V36h3.32v-4.65a3.28 3.28 0 0 0 1.61-2.8v0A3.27 3.27 0 0 0 24 25.28Z"/><rect width="33.23" height="25.73" x="7.38" y="17.77" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.32"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.35 17.77v-2.61a10.66 10.66 0 0 1 21.32 0v2.61"/>`),
		g.Group(children),
	)
}