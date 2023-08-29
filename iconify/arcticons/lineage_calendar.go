package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineageCalendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.604 5.5h18.793v4.495a3 3 0 0 1-3 3H17.604a3 3 0 0 1-3-3V5.5h0Zm0 3.748h18.793m-8.266 24.209a5.043 5.043 0 0 0 10.086 0v-5.138a5.043 5.043 0 0 0-10.086 0Zm-6.067-2.567a3.805 3.805 0 0 0 3.805-3.805h0a3.805 3.805 0 0 0-3.805-3.805m0 15.22a3.805 3.805 0 0 0 3.805-3.805h0a3.805 3.805 0 0 0-3.805-3.805m-6.28 6.326c1.051.88 2.186 1.284 4.734 1.284h1.546m-6.281-13.953c1.053-.878 2.189-1.278 4.737-1.271l1.545.003m-3.878 7.611h3.877"/>`),
		g.Group(children),
	)
}