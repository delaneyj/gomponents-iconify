package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Data(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 12v5c0 1.657-2.686 3-6 3s-6-1.343-6-3v-5m12 0V7m0 5c0 1.657-2.686 3-6 3s-6-1.343-6-3m12-5c0-1.657-2.686-3-6-3S6 5.343 6 7m12 0c0 1.657-2.686 3-6 3S6 8.657 6 7m0 5V7"/>`),
		g.Group(children),
	)
}