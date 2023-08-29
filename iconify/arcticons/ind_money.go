package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndMoney(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 18.929v11.183m3.267 0V18.929l7.409 11.183V11.888m2.915 18.224V18.93h2.516a4.893 4.893 0 0 1 4.893 4.893v1.398a4.893 4.893 0 0 1-4.893 4.892h-2.516Z"/><path d="m28.299 16.01l-4.123-4.122l-4.123 4.122"/></g>`),
		g.Group(children),
	)
}