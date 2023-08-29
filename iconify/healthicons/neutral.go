package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Neutral(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M24 40c8.837 0 16-7.163 16-16S32.837 8 24 8S8 15.163 8 24s7.163 16 16 16Zm0 2c9.941 0 18-8.059 18-18S33.941 6 24 6S6 14.059 6 24s8.059 18 18 18Z" clip-rule="evenodd"/><path d="M20 21c0 2.21-1.12 4-2.5 4S15 23.21 15 21s1.12-4 2.5-4s2.5 1.79 2.5 4Zm13 0c0 2.21-1.12 4-2.5 4S28 23.21 28 21s1.12-4 2.5-4s2.5 1.79 2.5 4Z"/><path fill-rule="evenodd" d="M16 31a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2H17a1 1 0 0 1-1-1Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}