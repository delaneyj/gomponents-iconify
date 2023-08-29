package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HammerAndAnvil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M6 14C6 4 14 4 14 4v20H6V14Zm8-4h28v6H14zM6 30h36s0 8-6 8h-7l2 6H13l2-6H6v-8Z"/>`),
		g.Group(children),
	)
}