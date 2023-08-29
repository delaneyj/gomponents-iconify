package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeBolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#fbc02d" d="M9.012 2h13.967q-2.787 5.6-5.593 11.194q2.8.014 5.6.009q-4.9 8.4-9.794 16.8c-.019-4.192-.009-8.375-.009-12.567c-1.391 0-2.782 0-4.173-.009Z"/>`),
		g.Group(children),
	)
}