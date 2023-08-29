package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeRiot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#ed1846" d="M2 2.914a.414.414 0 0 1 .415-.414H21a9 9 0 0 1 9 9v1.587a.414.414 0 0 1-.415.414h-5.17a.414.414 0 0 1-.415-.41V11.5a3 3 0 0 0-3-3H8.415A.414.414 0 0 0 8 8.917V29.1a.394.394 0 0 1-.416.4A6 6 0 0 1 2 23.51Z"/><path fill="#ed1846" d="M13.415 14.506a.394.394 0 0 0-.4.414A6 6 0 0 0 19 20.509h4a1 1 0 0 1 1 1V29.1a.394.394 0 0 0 .416.4A6 6 0 0 0 30 23.51v-2a7 7 0 0 0-7-7Z"/>`),
		g.Group(children),
	)
}