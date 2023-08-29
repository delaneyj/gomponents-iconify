package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignHorizontalLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 26H11a2.002 2.002 0 0 1-2-2v-4a2.002 2.002 0 0 1 2-2h15a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zm0-6.001L11 20v4h15zM18 14h-7a2.002 2.002 0 0 1-2-2V8a2.002 2.002 0 0 1 2-2h7a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zm0-6.001L11 8v4h7zM4 2h2v28H4z"/>`),
		g.Group(children),
	)
}