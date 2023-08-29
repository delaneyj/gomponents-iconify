package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExitToApp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M472 16H40a24.028 24.028 0 0 0-24 24v160h32V48h416v416H48V304H16v168a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V40a24.028 24.028 0 0 0-24-24Z"/><path fill="currentColor" d="m209.377 363.306l22.627 22.627L366.627 251.31L232.004 116.687l-22.627 22.626l95.997 95.998H16v32h289.372l-95.995 95.995z"/>`),
		g.Group(children),
	)
}