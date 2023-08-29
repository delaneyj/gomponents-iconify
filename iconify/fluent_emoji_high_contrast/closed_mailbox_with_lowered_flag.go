package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedMailboxWithLoweredFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M4.98 20h4.29c.55 0 .99.44.98.98c0 .54-.44.98-.98.98H4.98c-.54 0-.98-.44-.98-.98s.44-.98.98-.98Z"/><path d="M7.84 6A6.834 6.834 0 0 0 1 12.84v9.27a2.5 2.5 0 0 0 2.5 2.5h8.33v5.43h4.98v-5.43h9.18a2.5 2.5 0 0 0 2.5-2.5v-2.066h-2v2.066a.5.5 0 0 1-.5.5H14.67v-9.77A6.823 6.823 0 0 0 12.668 8h8.982c1.7 0 3.196.878 4.06 2.21h2.252A6.839 6.839 0 0 0 21.65 6H7.84Zm4.83 6.84v9.77H3.5a.5.5 0 0 1-.5-.5v-9.27A4.834 4.834 0 0 1 7.84 8a4.832 4.832 0 0 1 4.83 4.84Z"/><path d="M26 13.63h-5.484A2.42 2.42 0 0 1 16 12.42a2.42 2.42 0 0 1 4.516-1.21h8.904c.67 0 1.21.54 1.21 1.21c0 .46-.254.858-.63 1.063v4.648c0 .476-.369.869-.816.869h-2.369c-.446 0-.815-.393-.815-.869V13.63Z"/></g>`),
		g.Group(children),
	)
}