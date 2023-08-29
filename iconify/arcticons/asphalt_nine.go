package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AsphaltNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.345 4.5v31.244L23.95 43.5l15.704-7.472V4.5H8.345Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M15.432 11.816h17.136v10.977H15.432z"/><path d="M15.432 28.837v3.089l8.519 4.258l8.617-4.078v-9.314"/></g>`),
		g.Group(children),
	)
}