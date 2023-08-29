package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderArrowUpTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6.75A3.75 3.75 0 0 1 5.75 3h3.672c.729 0 1.428.29 1.944.805L13.25 5.69l-2.944 2.945A1.25 1.25 0 0 1 9.421 9H2V6.75Zm.004 3.75v9.75A3.75 3.75 0 0 0 5.754 24h8.745A7.5 7.5 0 0 1 26 14.401V9.75A3.75 3.75 0 0 0 22.25 6h-7.19l-3.694 3.695a2.75 2.75 0 0 1-1.944.805H2.004ZM27 19.5a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0Zm-7 4a.5.5 0 0 0 1 0v-6.793l2.146 2.147a.5.5 0 0 0 .708-.708l-3-3a.499.499 0 0 0-.351-.146h-.006a.5.5 0 0 0-.35.146l-3 3a.5.5 0 0 0 .707.708L20 16.707V23.5Z"/>`),
		g.Group(children),
	)
}