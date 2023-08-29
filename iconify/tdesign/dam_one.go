package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DamOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3c3.905 0 7 3.162 7 7.1V20h-2v-9.9a5 5 0 0 0-10 0V20H5v-9.9C5 6.162 8.095 3 12 3Zm9 17v-9.9C21 5.09 17.042 1 12 1s-9 4.09-9 9.1V20H2v2h8v-2H9v-9.9a3 3 0 1 1 6 0V20h-1v2h8v-2h-1Z"/>`),
		g.Group(children),
	)
}