package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StethoscopeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 8C18.1 8 17 9.1 17 10.5c0 1.2.9 2.2 2 2.4v2.6c0 3-2.5 5.5-5.5 5.5S8 18.5 8 15.5v-1.8l3.3-2.6C12.4 10.2 13 9 13 7.6V2.5c0-.3-.2-.5-.5-.5h-2c-.3 0-.5.2-.5.5s.2.5.5.5H12v4.6c0 1.1-.5 2.1-1.3 2.7l-3.2 2.5l-3.2-2.5C3.5 9.6 3 8.6 3 7.6V3h1.5c.3 0 .5-.2.5-.5S4.8 2 4.5 2h-2c-.3 0-.5.2-.5.5v5.1c0 1.4.6 2.7 1.7 3.5L7 13.7v1.8c0 3.6 2.9 6.5 6.5 6.5s6.5-2.9 6.5-6.5v-2.6c1.1-.2 2-1.2 2-2.4C22 9.1 20.9 8 19.5 8zm0 4c-.8 0-1.5-.7-1.5-1.5S18.7 9 19.5 9s1.5.7 1.5 1.5s-.7 1.5-1.5 1.5z"/>`),
		g.Group(children),
	)
}