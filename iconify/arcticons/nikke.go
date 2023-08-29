package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nikke(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.776 17.141v13.718m-8.276 0V17.141l5.961 13.718V17.141m4.568 0v13.718m4.836 0L21.16 24l3.705-6.813M21.16 24h-1.131m6.993-6.859v13.718m4.837 0L28.154 24l3.705-6.813M28.154 24h-1.132m6.979 0h2.933m1.566 6.859h-4.499V17.141H38.5"/>`),
		g.Group(children),
	)
}