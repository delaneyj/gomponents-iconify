package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentBlockSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 7.5c0 .972.308 1.872.832 2.607A4.48 4.48 0 0 0 16.5 12A4.5 4.5 0 1 0 12 7.5Zm4.5 3a2.985 2.985 0 0 1-1.524-.415l4.109-4.109A3 3 0 0 1 16.5 10.5Zm-2.585-1.476l4.109-4.109a3 3 0 0 0-4.109 4.109Z" clip-rule="evenodd"/><path fill="currentColor" d="M4.823 15.21C2.796 10.233 6.057 4.82 11.17 4.085c.17-.025.282.166.194.312a6 6 0 0 0 8.782 7.869c.139-.107.345-.01.333.165a8.733 8.733 0 0 1-8.711 8.119h-7.82a.5.5 0 0 1-.314-.89l1.972-1.583a.5.5 0 0 0 .15-.579l-.933-2.288Z"/>`),
		g.Group(children),
	)
}