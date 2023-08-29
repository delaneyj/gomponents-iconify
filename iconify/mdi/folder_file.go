package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 8c-2.21 0-4 1.79-4 4v8H4a2 2 0 0 1-2-2V6c0-1.11.89-2 2-2h6l2 2h8a2 2 0 0 1 2 2v2.17l-1.59-1.58l-.58-.59H15m8 6v7c0 1.11-.89 2-2 2h-6a2 2 0 0 1-2-2v-9c0-1.1.9-2 2-2h4l4 4m-2 .83L18.17 12H18v3h3v-.17Z"/>`),
		g.Group(children),
	)
}