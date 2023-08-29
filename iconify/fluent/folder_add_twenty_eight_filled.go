package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderAddTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6.75A3.75 3.75 0 0 1 5.75 3h3.672c.729 0 1.428.29 1.944.805L13.25 5.69l-2.944 2.945A1.25 1.25 0 0 1 9.421 9H2V6.75Zm.004 3.75v9.75A3.75 3.75 0 0 0 5.754 24h8.745A7.5 7.5 0 0 1 26 14.401V9.75A3.75 3.75 0 0 0 22.25 6h-7.19l-3.694 3.695a2.75 2.75 0 0 1-1.944.805H2.004ZM27 19.5a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0Zm-6-4a.5.5 0 0 0-1 0V19h-3.5a.5.5 0 0 0 0 1H20v3.5a.5.5 0 0 0 1 0V20h3.5a.5.5 0 0 0 0-1H21v-3.5Z"/>`),
		g.Group(children),
	)
}