package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeVhdl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#0d9b35" d="M2 2h28v28H2Zm1.689.067a1.624 1.624 0 0 0-1.626 1.625v24.622a1.625 1.625 0 0 0 1.626 1.626h24.627a1.625 1.625 0 0 0 1.626-1.626V3.692a1.624 1.624 0 0 0-1.626-1.625Zm-.681 26.012a.911.911 0 0 0 .911.912h24.164a.911.911 0 0 0 .911-.912V3.919a.91.91 0 0 0-.911-.911H3.919a.91.91 0 0 0-.911.911Z"/><path fill="#fff" d="m25.52 5.502l-6.662 20.989h-5.704L6.492 5.502h4.016l5.521 17.293l5.475-17.293h4.016z"/>`),
		g.Group(children),
	)
}