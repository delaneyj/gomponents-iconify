package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zahler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 29.5v-22a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2h0v22M19 16.723L24 14m0 0v20m9.5 3h7m-33 0h7M11 33.5v7"/><rect width="15" height="15" x="3.5" y="29.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><rect width="15" height="15" x="29.5" y="29.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.5 42.5h-11"/>`),
		g.Group(children),
	)
}