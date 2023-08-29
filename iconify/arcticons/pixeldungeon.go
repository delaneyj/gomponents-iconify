package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixeldungeon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41 5H7a2 2 0 0 0-2 2v34a2 2 0 0 0 2 2h34a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2Zm-2 17h-7.92v-5.08H26V9h13Zm-11.92-1.08v6.16h-6.16v-6.16ZM22 9v7.92h-5.08V22H9V9Zm15.12 30H9V26h7.92v5.08h14.16V26H39v11.12A1.88 1.88 0 0 1 37.12 39Z"/>`),
		g.Group(children),
	)
}