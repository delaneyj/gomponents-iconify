package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LiteralRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 17h1v2h-1c-.768 0-1.47-.289-2-.764A2.989 2.989 0 0 1 10 19H9v-2h1a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H9V1h1c.768 0 1.47.289 2 .764A2.989 2.989 0 0 1 14 1h1v2h-1a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1ZM8 5H2a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h6v-2H2V7h6V5Zm8 2V5h2a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-2v-2h2V7h-2Z"/>`),
		g.Group(children),
	)
}