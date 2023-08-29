package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HelpTriangleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M11.94 2a2.99 2.99 0 0 1 2.45 1.279l.108.164l8.431 14.074a2.989 2.989 0 0 1-2.366 4.474l-.2.009H3.507a2.99 2.99 0 0 1-2.648-4.308l.101-.189L9.385 3.438A2.989 2.989 0 0 1 11.94 2zM12 16a1 1 0 0 0-.993.883L11 17l.007.127a1 1 0 0 0 1.986 0L13 17.01l-.007-.127A1 1 0 0 0 12 16zm1.368-6.673a2.98 2.98 0 0 0-3.631.728a1 1 0 0 0 1.44 1.383l.171-.18a.98.98 0 0 1 1.11-.15a1 1 0 0 1-.34 1.886l-.232.012A1 1 0 0 0 11.997 15a3 3 0 0 0 1.371-5.673z"/></g>`),
		g.Group(children),
	)
}