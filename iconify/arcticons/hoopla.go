package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hoopla(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.501 18.372v8.642M8.5 23.45a2.126 2.126 0 1 1 4.252 0v3.564m4.074.012a2.15 2.15 0 0 1-2.125-2.16V23.46a2.126 2.126 0 1 1 4.251 0v1.404a2.15 2.15 0 0 1-2.126 2.16Zm6.2-.002a2.15 2.15 0 0 1-2.126-2.16v-1.405a2.126 2.126 0 1 1 4.251 0v1.404a2.15 2.15 0 0 1-2.125 2.161Zm4.074-2.166a2.126 2.126 0 1 0 4.252 0v-1.405a2.126 2.126 0 1 0-4.252 0m0-2.16v8.642m6.2-11.545v8.643m6.2-2.163a2.126 2.126 0 1 1-4.25 0v-1.405a2.126 2.126 0 1 1 4.25 0m.001 3.565v-5.725"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/>`),
		g.Group(children),
	)
}