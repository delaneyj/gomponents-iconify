package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kinoteatr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.07 20.64v9.2m-6.18 0v-9.2m6.18 0l-6.18 9.2M39 23.71a3.08 3.08 0 0 0-3.26-3.08a3.2 3.2 0 0 0-2.92 3.26v2.86a3.09 3.09 0 0 0 3.09 3.09h0A3.09 3.09 0 0 0 39 26.75v-3M9 20.61v9.22m1.19-4.61l3.41-4.58m-3.41 4.58l3.41 4.63m-3.41-4.63H9m15.35-4.58v9.2m6.18-9.2v9.2m-6.18-4.62h6.18"/><circle cx="35.91" cy="18.82" r=".66" fill="currentColor"/><circle cx="38.34" cy="18.82" r=".66" fill="currentColor"/><circle cx="33.49" cy="18.82" r=".66" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}