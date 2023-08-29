package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Castle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m4 11l7-7l7 7H4Zm26 0l7-7l7 7H30Zm14 33V26h-4v-9h-6v9h-6v-5l-4-4l-4 4v5h-6v-9H8v9H4v18h14v-4a6 6 0 0 1 12 0v4h14ZM7 11h8v6H7zm26 0h8v6h-8z"/>`),
		g.Group(children),
	)
}