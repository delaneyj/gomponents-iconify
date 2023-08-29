package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeBazelVersion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#76d275" d="m9 2l7 7l-7 7l-7-7z"/><path fill="#43a047" d="M2 9v7l7 7v-7z"/><path fill="#76d275" d="m23 2l7 7l-7 7l-7-7z"/><path fill="#43a047" d="M30 9v7l-7 7v-7zM16 9l7 7l-7 7l-7-7z"/><path fill="#00701a" d="M16 23v7l-7-7v-7z"/><path fill="#004300" d="m16 23l7-7v7l-7 7z"/><path fill="#f60" d="M9 2L2 9v7l14 14l14-14V9l-7-7l-7 7z"/>`),
		g.Group(children),
	)
}