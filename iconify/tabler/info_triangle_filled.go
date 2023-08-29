package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoTriangleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M11.94 2a2.99 2.99 0 0 1 2.45 1.279l.108.164l8.431 14.074a2.989 2.989 0 0 1-2.366 4.474l-.2.009H3.507a2.99 2.99 0 0 1-2.648-4.308l.101-.189L9.385 3.438A2.989 2.989 0 0 1 11.94 2zM12 12h-1l-.117.007a1 1 0 0 0 0 1.986L11 14v3l.007.117a1 1 0 0 0 .876.876L12 18h1l.117-.007a1 1 0 0 0 .876-.876L14 17l-.007-.117a1 1 0 0 0-.764-.857l-.112-.02L13 16v-3l-.007-.117a1 1 0 0 0-.876-.876L12 12zm.01-3l-.127.007a1 1 0 0 0 0 1.986L12 11l.127-.007a1 1 0 0 0 0-1.986L12.01 9z"/></g>`),
		g.Group(children),
	)
}