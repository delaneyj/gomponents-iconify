package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeWgsl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#005a9c" fill-rule="evenodd" d="M293.627 508.5L52.873 91.5h481.51z"/><path fill="#0066b0" fill-rule="evenodd" d="m534.628 91.5l-120.5 208h241z"/><path fill="#0076cc" fill-rule="evenodd" d="m534.628 507.5l-120.5-208h241z"/><path fill="#0086e8" fill-rule="evenodd" d="m654.628 300.5l-60.5-104h121z"/><path fill="#0093ff" fill-rule="evenodd" d="m654.628 92.5l-60.5 104h121z"/>`),
		g.Group(children),
	)
}