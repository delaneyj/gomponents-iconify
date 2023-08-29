package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaCastOutline0"><g id="evaCastOutline1"><g id="evaCastOutline2" fill="currentColor"><path d="M18.4 3H5.6A2.7 2.7 0 0 0 3 5.78V7a1 1 0 0 0 2 0V5.78A.72.72 0 0 1 5.6 5h12.8a.72.72 0 0 1 .6.78v12.44a.72.72 0 0 1-.6.78H17a1 1 0 0 0 0 2h1.4a2.7 2.7 0 0 0 2.6-2.78V5.78A2.7 2.7 0 0 0 18.4 3ZM3.86 14A1 1 0 0 0 3 15.17a1 1 0 0 0 1.14.83a2.49 2.49 0 0 1 2.12.72a2.52 2.52 0 0 1 .51 2.84a1 1 0 0 0 .48 1.33a1.06 1.06 0 0 0 .42.09a1 1 0 0 0 .91-.58A4.52 4.52 0 0 0 3.86 14Z"/><path d="M3.86 10.08a1 1 0 0 0 .28 2a6 6 0 0 1 5.09 1.71a6 6 0 0 1 1.53 5.95a1 1 0 0 0 .68 1.26a.9.9 0 0 0 .28 0a1 1 0 0 0 1-.72a8 8 0 0 0-8.82-10.2Z"/><circle cx="4" cy="19" r="1"/></g></g></g>`),
		g.Group(children),
	)
}