package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowExportRtlTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.5 4a.5.5 0 0 0-.5.5v11a.5.5 0 0 0 1 0v-11a.5.5 0 0 0-.5-.5ZM15 10a.5.5 0 0 0-.5-.5H3.707l3.147-3.146a.5.5 0 0 0-.708-.708l-4 4a.498.498 0 0 0 0 .708l4 4a.5.5 0 0 0 .708-.707L3.707 10.5H14.5a.5.5 0 0 0 .5-.5Z"/>`),
		g.Group(children),
	)
}