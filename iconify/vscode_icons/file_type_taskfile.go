package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeTaskfile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#28bdb0" d="M16.002 16.256v13.73l-11.75-7.12V9.132z"/><path fill="#68d1c7" d="M16.002 16.256v13.73l11.747-7.12V9.132z"/><path fill="#94ded7" d="M16.002 16.256L4.252 9.133l11.75-7.12l11.747 7.12z"/>`),
		g.Group(children),
	)
}