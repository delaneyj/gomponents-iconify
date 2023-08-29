package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeDevcontainer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="16" cy="16" r="14" fill="#193e63"/><path fill="#add1ea" d="m10.777 22.742l-1.434-1.394l3.386-3.483l-3.383-3.448l1.428-1.4l4.751 4.842l-4.748 4.883zm10.643-3.641l1.434-1.395l-3.386-3.482l3.383-3.448l-1.428-1.4l-4.751 4.842l4.748 4.883z"/>`),
		g.Group(children),
	)
}