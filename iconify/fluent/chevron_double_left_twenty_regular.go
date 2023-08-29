package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleLeftTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.353 15.854a.5.5 0 0 1-.707.001L5.162 10.39a.55.55 0 0 1 0-.78l5.484-5.464a.5.5 0 1 1 .706.708L6.188 10l5.164 5.147a.5.5 0 0 1 .001.707Zm4 0a.5.5 0 0 1-.708.001L9.161 10.39a.55.55 0 0 1 0-.78l5.484-5.464a.5.5 0 1 1 .706.708L10.187 10l5.164 5.147a.5.5 0 0 1 .001.707Z"/>`),
		g.Group(children),
	)
}