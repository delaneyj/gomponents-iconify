package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeGenstat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#e2e2e2" d="M2 5.861h28v20H2z"/><path fill="#6cc1ee" d="M2 2h28v21.613c-9.333.915-12.5-17.32-17.956-17.32C7.455 6.293 5.7 16.57 2 17.787Z"/><path fill="#46b270" d="M2 30h28v-4.639C20.382 26.388 15.526 8.083 11.959 8.042c-3.167 0-4.9 10.382-9.959 11.494Z"/>`),
		g.Group(children),
	)
}