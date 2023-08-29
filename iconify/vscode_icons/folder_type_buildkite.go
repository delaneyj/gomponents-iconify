package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderTypeBuildkite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#55bf91" d="M27.5 5.5h-9.3l-2.1 4.2H4.4v16.8h25.2v-21Zm0 4.2h-8.2l1.1-2.1h7.1Z"/><path fill="#30f2a2" d="m7 15l8 3.92v8L7 23Zm16.08 0L31 18.92L23.08 23Z"/><path fill="#14cc80" d="M23.08 15L15 18.92v8L23.08 23ZM31 18.92L23.08 23v8L31 26.92Z"/>`),
		g.Group(children),
	)
}