package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StrikingSplinter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m23.018 20.705l135.64 163.623l-107.33-32.39l168.79 111.326L82.784 224.11l192.51 111.87l-130.525-1.76l282.08 126.116a428.486 428.486 0 0 0 42.728 19.246l2.297.885l20.797 9.3l-16.895-37.82a398.823 398.823 0 0 0-12.03-26.926L338.312 144.24l1.094 129.362L228.352 82.393l38.482 136.49L155.906 50.668l31.684 106.467L23.018 20.705zm225.148 225.178c94.262 38.75 169.608 116.195 208.152 207.924c-91.01-40.827-168.835-115.908-208.152-207.924z"/>`),
		g.Group(children),
	)
}