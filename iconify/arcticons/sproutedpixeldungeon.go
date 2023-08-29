package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sproutedpixeldungeon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41 5H7a2 2 0 0 0-2 2v34a2 2 0 0 0 2 2h34a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2Zm-2 17h-7.92v-5.08H26V9h13Zm-11.92-1.08v6.16h-6.16v-6.16ZM22 9v7.92h-5.08V22H9V9Zm15.12 30H9V26h7.92v5.08h14.16V26H39v11.12A1.88 1.88 0 0 1 37.12 39Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.9 16.75c-.28-2.32 2.71-5.28 7.53-5a5.26 5.26 0 0 1-1.43 4.7c-2.19 2.4-3.88 2-3.88 2m-12.07-1.7c.29-2.32-3.1-5.61-8.37-4.27a5.49 5.49 0 0 0 1.91 4.6c2.28 2.13 4.29 1.35 4.29 1.35M24 31.08c.08 1.37-.09 5.31-2.5 5.67"/>`),
		g.Group(children),
	)
}