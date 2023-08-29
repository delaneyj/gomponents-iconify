package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clubs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3a4 4 0 0 1 3.164 6.447A4 4 0 1 1 14 15.645V17l1 4H9l1-4v-1.355a4 4 0 1 1-1.164-6.199A4 4 0 0 1 11.999 3z"/>`),
		g.Group(children),
	)
}