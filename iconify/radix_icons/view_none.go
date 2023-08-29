package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewNone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 2.587L1.852 13H13.5a.5.5 0 0 0 .5-.5V2.587ZM.763 13.807l.062.073l.03-.026c.195.094.414.146.645.146h12a1.5 1.5 0 0 0 1.5-1.5v-10a1.5 1.5 0 0 0-.763-1.307l-.062-.073l-.03.025A1.494 1.494 0 0 0 13.5 1h-12A1.5 1.5 0 0 0 0 2.5v10c0 .56.307 1.05.763 1.307ZM1 12.413L13.148 2H1.5a.5.5 0 0 0-.5.5v9.913Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}