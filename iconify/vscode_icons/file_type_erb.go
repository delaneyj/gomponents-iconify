package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeErb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#921a1e" d="M7.25 24.75h17.5L12.5 12.5l-5.25 5.25ZM2 30h28V2H16L2 16Zm25.375-2.625H4.625v-10.5l12.25-12.25h10.5Z"/>`),
		g.Group(children),
	)
}