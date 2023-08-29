package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeMondoo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#6e45e2" d="M9.03 9.03A6.97 6.97 0 0 0 2.06 16a6.97 6.97 0 0 0 6.97 6.971V9.03zm0 6.97A6.97 6.97 0 0 0 16 22.971a6.97 6.97 0 0 0 6.971-6.97a6.97 6.97 0 0 0-6.97-6.972A6.97 6.97 0 0 0 9.028 16zm13.942 0v6.97A6.968 6.968 0 0 0 29.94 16a6.97 6.97 0 0 0-6.968-6.97V16z"/>`),
		g.Group(children),
	)
}