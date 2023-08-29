package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeCheader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#005f91" d="M8.329 29V3h3.192v9.329a7.132 7.132 0 0 1 5.64-2.589a7.605 7.605 0 0 1 3.636.825a4.842 4.842 0 0 1 2.208 2.279a10.506 10.506 0 0 1 .665 4.221V29h-3.192V17.064a4.932 4.932 0 0 0-1.038-3.485a3.858 3.858 0 0 0-2.935-1.091a5.176 5.176 0 0 0-2.669.736a4.157 4.157 0 0 0-1.782 2a9.164 9.164 0 0 0-.532 3.476V29Z"/>`),
		g.Group(children),
	)
}