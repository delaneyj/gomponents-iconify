package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleUpTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.146 9.646a.5.5 0 0 0 .708.706L10 5.188l5.146 5.164a.5.5 0 1 0 .708-.706L10.39 4.162a.55.55 0 0 0-.78 0L4.146 9.646Zm0 5.001a.5.5 0 0 0 .708.706L10 10.188l5.146 5.165a.5.5 0 1 0 .708-.706L10.39 9.163a.55.55 0 0 0-.78 0l-5.464 5.484Z"/>`),
		g.Group(children),
	)
}