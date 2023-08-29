package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Watch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feWatch0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feWatch1" fill="currentColor"><path id="feWatch2" d="M15 6.803A5.998 5.998 0 0 1 18 12c0 2.22-1.207 4.16-3 5.197V20a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2v-2.803A5.998 5.998 0 0 1 6 12c0-2.22 1.207-4.16 3-5.197V4a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2.803ZM12 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm2-5a1 1 0 0 1 0 2h-1c-1.1 0-2-.9-2-2v-1a1 1 0 0 1 2 0v1h1Z"/></g></g>`),
		g.Group(children),
	)
}