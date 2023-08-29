package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 10h-6v11H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v8h-2v-3Zm-8 0H5v4h6v-4Zm0 9v-3H5v3h6Zm2-14v3h6V5h-6Zm-2 0H5v3h6V5Z"/><path fill="currentColor" d="M16 18v-2h8v2h-8Z"/>`),
		g.Group(children),
	)
}