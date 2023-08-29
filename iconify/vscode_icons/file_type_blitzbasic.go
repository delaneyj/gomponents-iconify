package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeBlitzbasic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#00ffae" d="M4 9v3h1v7H4v3h10v-1h1v-1h1v-4h-1v-1h1v-4h-1v-1h-1V9H4Zm6 3h2v2h-2v-2Zm0 5h2v2h-2v-2Z"/><path fill="#00d8ff" d="M16 9v3h1v7h-1v3h10v-1h1v-1h1v-4h-1v-1h1v-4h-1v-1h-1V9H16Zm5 3h2v2h-2v-2Zm0 5h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}