package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldMinimalisticBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M3.378 5.082C3 5.62 3 7.22 3 10.417v1.574c0 5.638 4.239 8.375 6.899 9.536c.721.315 1.082.473 2.101.473V2c-.811 0-1.595.268-3.162.805L8.265 3c-3.007 1.03-4.51 1.545-4.887 2.082Z" clip-rule="evenodd"/><path d="M21 11.991v-1.574c0-3.198 0-4.797-.378-5.335c-.377-.537-1.88-1.052-4.887-2.081l-.573-.196C13.595 2.268 12.812 2 12 2v20c1.02 0 1.38-.158 2.101-.473C16.761 20.365 21 17.63 21 11.991Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}