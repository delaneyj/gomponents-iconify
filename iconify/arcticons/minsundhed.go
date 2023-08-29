package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minsundhed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.714 21.473c1.308.914.623 3.325-2.97 6.258c0 0-2.801-2.12-2.85-4.17s1.749-1.99 2.543.282c0 0 1.97-3.284 3.277-2.37Zm19.961 0c1.308.914.623 3.325-2.97 6.258c0 0-2.801-2.12-2.85-4.17s1.749-1.99 2.543.282c0 0 1.97-3.284 3.277-2.37Zm-7.77 12.528h-9.81s.477 5.439 4.947 5.439S28.905 34 28.905 34Z"/>`),
		g.Group(children),
	)
}