package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeStencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#7b79ff" d="M30 12.849H7.83L2 19.151h22.17Zm-15.217-7.47h10.751l-5.806 6.3H8.968Zm-2.373 14.94H23.2l-5.833 6.3H6.634Z"/>`),
		g.Group(children),
	)
}