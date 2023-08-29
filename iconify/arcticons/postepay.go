package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Postepay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.148 24.608a3.246 3.246 0 0 0 3.277 3.978h0a4.951 4.951 0 0 0 4.679-3.978l.456-2.586a3.246 3.246 0 0 0-3.277-3.977h0a4.951 4.951 0 0 0-4.679 3.977m.701-3.977L13.5 33.955m12.527-9.347a3.246 3.246 0 0 0 3.276 3.978h0a4.951 4.951 0 0 0 4.68-3.978l.455-2.586a3.246 3.246 0 0 0-3.276-3.977h0a4.951 4.951 0 0 0-4.68 3.977m.702-3.977l-2.806 15.91"/>`),
		g.Group(children),
	)
}