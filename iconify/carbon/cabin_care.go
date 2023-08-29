package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CabinCare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27.302 2a2.662 2.662 0 0 0-1.908.806l-.393.405l-.397-.405a2.661 2.661 0 0 0-3.816 0a2.8 2.8 0 0 0 0 3.896L25.001 11l4.209-4.298a2.8 2.8 0 0 0 0-3.896A2.661 2.661 0 0 0 27.302 2zm-3.48 16h-8.739L11.864 5.965a4 4 0 0 0-7.727 2.07l4.408 16.482A2 2 0 0 0 10.477 26H19v2H4v2h15a2 2 0 0 0 2-2v-2h3a4.005 4.005 0 0 0 3.98-4.4a4.121 4.121 0 0 0-4.158-3.6zM24 24H10.478L6.068 7.518a2 2 0 1 1 3.864-1.035L13.547 20H24a2 2 0 0 1 0 4z"/>`),
		g.Group(children),
	)
}