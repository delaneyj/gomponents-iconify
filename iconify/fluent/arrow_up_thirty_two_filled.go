package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 29c-.69 0-1.25-.56-1.25-1.25V7.213l-7.628 7.432a1.25 1.25 0 1 1-1.744-1.79l9.747-9.497a1.246 1.246 0 0 1 1.75 0l9.747 9.497a1.25 1.25 0 1 1-1.744 1.79L17.25 7.213V27.75c0 .69-.56 1.25-1.25 1.25Z"/>`),
		g.Group(children),
	)
}