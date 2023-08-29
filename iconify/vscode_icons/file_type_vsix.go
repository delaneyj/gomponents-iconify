package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeVsix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#cfcfcf" d="M2 16v14h13v-6H8V8h7V2H2ZM17 5v3h4V6h5v5h-2v4h6V2H17Z"/><path fill="#cfcfcf" d="M10 16v6h12V10H10Z"/><path fill="#cfcfcf" d="M24 20.5V24h-7v6h13V17h-6Z"/>`),
		g.Group(children),
	)
}