package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsergroupAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 5a3 3 0 0 0 0 6v2a5 5 0 0 0-5 5v4H2v-4a7 7 0 0 1 3.75-6.2A5 5 0 0 1 9 3h1v2H9Zm6 0a3 3 0 1 0 0 6a3 3 0 0 0 0-6Zm-5 3a5 5 0 1 1 10 0a5 5 0 0 1-10 0ZM7 19a5 5 0 0 1 5-5h3v2h-3a3 3 0 0 0-3 3v1h6v2H7v-3Zm14-5v3h3v2h-3v3h-2v-3h-3v-2h3v-3h2Z"/>`),
		g.Group(children),
	)
}