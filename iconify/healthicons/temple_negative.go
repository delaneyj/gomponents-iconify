package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TempleNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsTempleNegative0)"><path d="M13 36.048v6h2v-6h-2Zm6 6h-2v-6h14v6h-2v-2a2 2 0 0 0-2-2h-6a2 2 0 0 0-2 2v2Zm14 0v-6h2v6h-2Zm-6-24h2v-4H19v4h2v-1a1 1 0 1 1 2 0v1h2v-1a1 1 0 1 1 2 0v1Zm-10 7v4h2v-1a1 1 0 1 1 2 0v1h2v-1a1 1 0 1 1 2 0v1h2v-1a1 1 0 1 1 2 0v1h2v-4H17Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM17 14.048s-3 .286-3-3.286c4 4.286 10-6.714 10-6.714s6 11 10 6.714c0 3.572-3 3.286-3 3.286v4.632c.962 1.649 3.728 5.781 6 3.257c0 3.11-2.8 3.11-2.8 3.11H33v4.948c.94 1.797 3.251 5.317 6 2.942c0 1.67-1.038 2.444-2 2.802v8.309H11v-8.309c-.961-.358-2-1.131-2-2.802c3.6 3.11 6-3.89 6-3.89v-4h-1.2s-2.8 0-2.8-3.11c2.8 3.11 6-3.89 6-3.89v-4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsTempleNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}