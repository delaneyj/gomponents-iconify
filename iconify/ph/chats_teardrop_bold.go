package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatsTeardropBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M172.29 68.9A84 84 0 0 0 12 104v66a18 18 0 0 0 18 18h54.1a84.18 84.18 0 0 0 75.9 48h66a18 18 0 0 0 18-18v-66a84 84 0 0 0-71.71-83.1ZM36 104a60 60 0 1 1 60 60H36Zm184 108h-60a60.14 60.14 0 0 1-49-25.37a83.93 83.93 0 0 0 68.55-91.37A60 60 0 0 1 220 152Z"/>`),
		g.Group(children),
	)
}