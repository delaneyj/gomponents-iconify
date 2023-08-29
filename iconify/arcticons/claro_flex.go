package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClaroFlex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="4.202" height="5.568" x="28.632" y="24.049" rx="2.101" ry="2.101"/><path d="M25.142 26.15a2.1 2.1 0 0 1 2.101-2.101h0m-2.101 0v5.568m-8.097-8.405v8.405m6.107-2.101a2.1 2.1 0 0 1-2.1 2.1h0a2.1 2.1 0 0 1-2.102-2.1V26.15a2.1 2.1 0 0 1 2.101-2.101h0c1.16 0 2.102.94 2.102 2.1m-.001 3.468v-5.568m-8.084 2.749v.035a2.784 2.784 0 0 1-2.784 2.784h0A2.784 2.784 0 0 1 9.5 26.833v-2.837a2.784 2.784 0 0 1 2.784-2.784h0a2.784 2.784 0 0 1 2.784 2.784v.034m15.665-2.083v-3.564m3.168 4.941l3.152-3.152m-2.117 6.661H38.5"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m38.5 34.722l-3.499 4.637m3.499 0l-3.499-4.637m-1.276 3.754a1.75 1.75 0 0 1-1.52.883h0a1.75 1.75 0 0 1-1.75-1.75v-1.137c0-.966.783-1.75 1.75-1.75h0c.966 0 1.75.784 1.75 1.75v.569h-3.5m-6.003-2.319h2.335m-1.412 4.638v-5.776c0-.677.548-1.225 1.225-1.225h.476c.603 0 .988.179 1.246.512v5.612c0 .483.392.875.875.875h.262"/>`),
		g.Group(children),
	)
}