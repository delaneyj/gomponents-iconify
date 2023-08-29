package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeBallerina(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#20b4ae" d="M8 9.859V2h6.818v10.376Zm0 2.461l4.666 1.723L8 15.764Zm6.818 3.389v3.805L11.5 30H8V18.225ZM24 9.859V2h-6.819v10.376Zm0 2.461l-4.668 1.723L24 15.764Zm-6.819 3.389v3.805L20.5 30H24V18.225Z"/>`),
		g.Group(children),
	)
}