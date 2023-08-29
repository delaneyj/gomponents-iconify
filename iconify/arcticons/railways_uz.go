package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RailwaysUz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m39.304 23.013l-1.728-3.456c-.987-1.975-1.878-2.635-3.95-2.962h-9.192c1.553 1.758 1.854 3.848.8 6.418h14.07Zm.987 1.974l3.209 6.418H24.247c1.776-1.903 1.874-4.08.987-6.418h15.057Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 16.595l4.937 10.367l5.43-10.367h4.937c5.17 0 4.937 7.405 0 7.405c4.937 0 5.43 7.405-.987 7.405h-5.925m-3.455-4.443l-2.312 4.443"/>`),
		g.Group(children),
	)
}