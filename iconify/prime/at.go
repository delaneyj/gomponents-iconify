package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func At(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3.25A8.75 8.75 0 0 0 3.25 12A8.65 8.65 0 0 0 12 20.75a.75.75 0 0 0 0-1.5A7.17 7.17 0 0 1 4.75 12A7.26 7.26 0 0 1 12 4.75c4.81 0 7.25 2.44 7.25 7.25v1.38a1.46 1.46 0 1 1-2.91 0v-5a.75.75 0 0 0-1.5 0v.34A4.31 4.31 0 0 0 12 7.66a4.34 4.34 0 0 0 0 8.68a4.3 4.3 0 0 0 3.24-1.49a2.95 2.95 0 0 0 5.51-1.47V12c0-5.64-3.11-8.75-8.75-8.75Zm0 11.59A2.84 2.84 0 1 1 14.84 12A2.85 2.85 0 0 1 12 14.84Z"/>`),
		g.Group(children),
	)
}