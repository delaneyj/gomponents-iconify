package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeCppTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#005f91" d="M26.914 13.8v1.54h-1.542v1.543h1.542v1.543h1.543v-1.543H30V15.34h-1.543V13.8Zm-3.5 0H21.87v1.54h-1.543v1.543h1.543v1.543h1.543v-1.543h1.543V15.34h-1.543Zm-3.654 5.226a6.167 6.167 0 1 1-.04-6.118l5.39-3.1a12.384 12.384 0 1 0 .032 12.34Z"/>`),
		g.Group(children),
	)
}