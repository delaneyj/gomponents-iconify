package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeSpacengine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#4780f5" d="M24 5h2v22h-2zm-5 4.409L6 13.941v2.13l13 5v-2.142L8.906 15.046L19 11.527V9.409z"/>`),
		g.Group(children),
	)
}