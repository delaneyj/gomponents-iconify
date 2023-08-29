package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeathAltTwoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M17.382 6.001L13.07 17.184l4.574 24.8l12.673.014l4.614-24.796l-4.294-11.19l-13.254-.01Zm14.629-1.988L16.01 4L11 16.992l4.978 26.99l16 .018L37 17.013l-4.99-13Z"/><path d="M23 17.997V26h2v-8.003h3v-2h-3V13h-2v2.997h-3v2h3Z"/></g>`),
		g.Group(children),
	)
}