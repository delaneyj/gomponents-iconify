package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderFileOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18h7v2H4a2 2 0 0 1-2-2V6c0-1.11.89-2 2-2h6l2 2h8a2 2 0 0 1 2 2v2.17l-1.59-1.58l-.41-.42V8H4v10m19-4v7c0 1.11-.89 2-2 2h-6a2 2 0 0 1-2-2v-9c0-1.1.9-2 2-2h4l4 4m-2 1h-3v-3h-3v9h6v-6Z"/>`),
		g.Group(children),
	)
}