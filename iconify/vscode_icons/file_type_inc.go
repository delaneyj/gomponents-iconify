package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeInc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="gray" d="M26 29H6V3h14l6 6v20z"/><path fill="#c5c5c5" d="M20.414 2H5v28h22V8.586ZM7 28V4h12v6h6v18Z"/><path fill="#c5c5c5" d="M2 15v5h9v4l7-7l-7-6v4H2z"/>`),
		g.Group(children),
	)
}