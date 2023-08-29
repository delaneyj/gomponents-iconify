package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeServerless(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#fd5750" fill-rule="evenodd" d="M2 22.419h4.956L5.42 27H2Zm0-8.709h7.875L8.34 18.29H2ZM2 5h10.794l-1.535 4.581H2Zm15.679 0H30v4.581H16.143Zm-4.455 13.291l1.536-4.581H30v4.581Zm-1.383 4.128H30V27H10.305Z"/>`),
		g.Group(children),
	)
}