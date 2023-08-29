package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoinCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.787 35.909v-8h1.3a4.012 4.012 0 0 1 4 4h0a4.012 4.012 0 0 1-4 4Zm-7.583 0v-8h2.6a2.7 2.7 0 0 1 0 5.4h-2.6m2.79-.007l2.51 2.607m-8.602-2.7h-3.4m4.3 2.7l-2.6-8l-2.7 8m.9-2.7h3.5m-6.512 0h0a2.689 2.689 0 0 1-2.678 2.7h-.022a2.689 2.689 0 0 1-2.7-2.678V30.51a2.689 2.689 0 0 1 2.677-2.7h.023a2.606 2.606 0 0 1 2.602 2.61c0 .03 0 .06-.002.09h0m14.281-18.268v8m-19.582 0v-8l4 8l4-8v8m14.722 0v-8l5.3 8v-8m-16.881 5.3a2.65 2.65 0 0 0 5.3.1q.002-.05 0-.1v-2.7a2.689 2.689 0 0 0-2.677-2.7h-.022a2.606 2.606 0 0 0-2.602 2.61l.002.09Z"/>`),
		g.Group(children),
	)
}