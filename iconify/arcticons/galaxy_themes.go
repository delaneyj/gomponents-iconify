package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GalaxyThemes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.966 14.615V6.04a25.131 25.131 0 0 0 11.343 2.768c6.482 0 8.575-3.308 15.799-3.308a22.973 22.973 0 0 1 6.926 1.507v7.608m0 0v11.546c0 3.375-3.916 3.646-7.494 3.646s-6.538 1.508-6.538 4.895c0 1.528 1.001 1.856 1.001 4.22S26.748 42.5 24 42.5s-4.003-1.215-4.003-3.579s1.001-2.69 1.001-4.22c0-3.387-2.96-4.894-6.538-4.894s-7.494-.27-7.494-3.646V14.615m0 0h34.068"/>`),
		g.Group(children),
	)
}