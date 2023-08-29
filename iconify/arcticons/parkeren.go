package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Parkeren(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.843 27.35V14.806h4.077a4.234 4.234 0 1 1 0 8.467h-4.077"/><circle cx="24" cy="21.078" r="11.729" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.596 43.301a.994.994 0 0 1-1.192 0C18.46 39.617 7.899 30.866 7.899 20.92a16.105 16.105 0 1 1 32.203 0c0 9.946-10.562 18.697-15.506 22.381Z"/>`),
		g.Group(children),
	)
}