package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BreastPump(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M35 25c0-4-6-6-6-6v-5H17v5s-6 2-6 6v19h24V25ZM20 4l-7 6m10 4l-6-7m9 1h9v7l6 5"/>`),
		g.Group(children),
	)
}