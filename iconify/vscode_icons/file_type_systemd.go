package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeSystemd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#d2d2d2" d="M2 12v8h3.256v-1.231H3.3v-5.538h1.956V12Zm24.744 0v1.231H28.7v5.538h-1.956V20H30v-8Z"/><path fill="#30d475" d="m17.628 16l5.21-2.769v5.538Z"/><ellipse cx="12.093" cy="16" fill="#30d475" rx="2.93" ry="2.769"/>`),
		g.Group(children),
	)
}