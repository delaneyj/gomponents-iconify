package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.001 1h12.414l5.586 5.587V13h-2V9h-6V3h-8v18h6v2h-8V1Zm12 2.415V7h3.586L15 3.415Zm8.663 11.841l-2.776 2.748l2.776 2.749l-1.407 1.421l-4.212-4.17l4.212-4.17l1.407 1.422ZM13 21h4.5v2H13v-2Z"/>`),
		g.Group(children),
	)
}