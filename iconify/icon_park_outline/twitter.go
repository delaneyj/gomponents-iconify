package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twitter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M5 35.762c1.929 1.067 15.891 9.115 25.82 2.912c9.928-6.203 9.38-16.89 9.38-21.788c.9-1.884 2.8-2.842 2.8-7.942c-1.866 1.724-3.721 2.31-5.565 1.76c-1.806-2.754-4.291-3.973-7.456-3.655c-4.746.477-6.482 5.133-5.971 11.158c-7.318 3.7-13.056-2.683-16.014-7.503c-.988 3.796-1.94 8.354 0 13.395c1.294 3.362 4.405 6.203 9.331 8.526C12.332 35.33 8.224 36.377 5 35.762Z"/>`),
		g.Group(children),
	)
}