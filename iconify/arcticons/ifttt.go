package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ifttt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 20.9v6.2m9.1-6.2h4.1m-2.1 6.2v-6.2M14.8 24h2m-2 3.1v-6.2h3m8.6 0h4.1m-2 6.2v-6.2m4.4 0H37m-2.1 6.2v-6.2"/>`),
		g.Group(children),
	)
}