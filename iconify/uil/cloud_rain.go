package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudRain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.56 12.67a1 1 0 0 0-1.12 0c-.11.08-2.69 1.86-2.69 4.58a3.25 3.25 0 0 0 6.5 0c0-2.75-2.58-4.51-2.69-4.58ZM12 18.5a1.25 1.25 0 0 1-1.25-1.25A3.66 3.66 0 0 1 12 14.8a3.61 3.61 0 0 1 1.25 2.45A1.25 1.25 0 0 1 12 18.5Zm6.42-10.78A7 7 0 0 0 5.06 9.61a4 4 0 0 0 .61 7.87h.08a1 1 0 0 0 1-.92a1 1 0 0 0-.92-1.08A2 2 0 0 1 4 13.5a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67a3 3 0 0 1 .43 5.79a1 1 0 0 0 .62 1.9a5 5 0 0 0-.14-9.53Z"/>`),
		g.Group(children),
	)
}