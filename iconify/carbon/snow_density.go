package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SnowDensity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 28a9.011 9.011 0 0 1-9-9a9.984 9.984 0 0 1 1.494-4.955l6.659-10.608a1.04 1.04 0 0 1 1.694 0l6.63 10.556A10.063 10.063 0 0 1 25 19a9.011 9.011 0 0 1-9 9Zm0-22.152l-5.782 9.208A7.977 7.977 0 0 0 9 19a7 7 0 0 0 14 0a8.062 8.062 0 0 0-1.248-3.995Z"/><path fill="currentColor" d="m20.6 16.8l-1.2-1.6L17 17v-3h-2v3l-2.4-1.8l-1.2 1.6l2.933 2.2l-2.933 2.2l1.2 1.6L15 21v3h2v-3l2.4 1.8l1.2-1.6l-2.933-2.2l2.933-2.2z"/>`),
		g.Group(children),
	)
}