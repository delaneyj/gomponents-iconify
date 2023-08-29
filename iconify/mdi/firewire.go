package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Firewire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 10c1.1 0 2 .9 2 2s-.9 2-2 2s-2-.9-2-2s.9-2 2-2m0-2c-2.2 0-4 1.8-4 4s1.8 4 4 4s4-1.8 4-4s-1.8-4-4-4m3 9H9v2h6v-2m0 3H9v2h6v-2m-8.7-9.4l4.2-4.2l-1.3-1.5l-4.3 4.3l1.4 1.4M4.2 8.5l4.2-4.2l-1.3-1.5l-4.3 4.3l1.4 1.4M15 17H9v2h6v-2M13.4 6.3l4.2 4.2L21.1 7l-4.2-4.2l-3.5 3.5Z"/>`),
		g.Group(children),
	)
}