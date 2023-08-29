package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Simpill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.25 14.823h11.5a4 4 0 0 1 4 4V41.5a2 2 0 0 1-2 2h-15.5a2 2 0 0 1-2-2V18.823a4 4 0 0 1 4-4Z"/><rect width="16.059" height="8.029" x="15.971" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.985 14.823v-2.294m8.03 2.294v-2.294m-10.324 8.603h12.618v16.059H17.691zm2.868 8.03h6.882m-6.882 3.441h6.882m-6.882-6.882h2.294m2.294 0h2.294"/>`),
		g.Group(children),
	)
}