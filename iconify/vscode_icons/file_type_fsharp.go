package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeFsharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#378bba" d="M2 16L15.288 2.712v6.644L8.644 16l6.644 6.644v6.644Z"/><path fill="#378bba" d="m10.542 16l4.746-4.746v9.492Z"/><path fill="#30b9db" d="M30 16L16.237 2.712v6.644L22.881 16l-6.644 6.644v6.644Z"/>`),
		g.Group(children),
	)
}