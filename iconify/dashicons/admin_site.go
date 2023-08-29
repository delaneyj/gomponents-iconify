package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdminSite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 0a9 9 0 1 0 0 18A9 9 0 0 0 9 0Zm3.46 11.95c0 1.47-.8 3.3-4.06 4.7c.3-4.17-2.52-3.69-3.2-5A3.25 3.25 0 0 1 7 9.1a8.49 8.49 0 0 1-4.18-2c.05.47.279.904.64 1.21a4.18 4.18 0 0 1-1.94-1.5a7.94 7.94 0 0 1 7.25-5.63c-.84 1.38-1.5 4.13 0 5.57C7.23 7 6.26 5 5.41 5.79c-1.13 1.06.33 2.51 3.42 3.08c3.29.59 3.66 1.58 3.63 3.08Zm1.34-4c-.32-1.11.62-2.23 1.69-3.14a7.27 7.27 0 0 1 .84 6.68c-.77-1.89-2.17-2.32-2.53-3.57v.03Z"/>`),
		g.Group(children),
	)
}