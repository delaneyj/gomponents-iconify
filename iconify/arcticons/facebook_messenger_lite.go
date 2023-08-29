package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacebookMessengerLite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.324 33.435A21.51 21.51 0 0 0 24 2.5h0A21.47 21.47 0 0 0 5.045 34.109l-.56 2.084l-1.928 7.202a1.673 1.673 0 0 0 2.049 2.048l7.2-1.928l2.074-.555a21.484 21.484 0 0 0 19.553.366M20.524 17.039l6.668 6.667l9.776-6.667l-6.667 9.776l-2.825 4.146l-6.667-6.667l-9.777 6.667l6.668-9.776Z"/><circle cx="38.5" cy="38.5" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.5 34.5v8h4"/>`),
		g.Group(children),
	)
}