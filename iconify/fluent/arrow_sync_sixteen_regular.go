package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSyncSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.146.646a.5.5 0 0 1 .708 0l1.5 1.5a.5.5 0 0 1 0 .708l-1.5 1.5a.5.5 0 1 1-.708-.708l.643-.642a5 5 0 0 0-2.937 8.88a.5.5 0 1 1-.63.777A6 6 0 0 1 7.797 2.003l-.65-.65a.5.5 0 0 1 0-.707Zm3.929 2.766a.5.5 0 0 1 .703-.073a6 6 0 0 1-3.575 10.658l.65.65a.5.5 0 0 1-.707.707l-1.5-1.5a.5.5 0 0 1 0-.708l1.5-1.5a.5.5 0 0 1 .708.708l-.643.642a5 5 0 0 0 2.937-8.88a.5.5 0 0 1-.073-.704Z"/>`),
		g.Group(children),
	)
}