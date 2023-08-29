package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderArrowRightTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6.75A3.75 3.75 0 0 1 5.75 3h3.672c.729 0 1.428.29 1.944.805L13.25 5.69l-2.944 2.945A1.25 1.25 0 0 1 9.421 9H2V6.75Zm.004 3.75v9.75A3.75 3.75 0 0 0 5.754 24h8.745A7.5 7.5 0 0 1 26 14.401V9.75A3.75 3.75 0 0 0 22.25 6h-7.19l-3.694 3.695a2.75 2.75 0 0 1-1.944.805H2.004ZM20.5 13a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13Zm-4 6a.5.5 0 0 0 0 1h6.793l-2.147 2.146a.5.5 0 0 0 .708.708l3-3a.5.5 0 0 0 .146-.351v-.006a.499.499 0 0 0-.146-.35l-3-3a.5.5 0 0 0-.708.707L23.293 19H16.5Z"/>`),
		g.Group(children),
	)
}