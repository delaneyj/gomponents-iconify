package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeDiff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#c00000" d="M6.975 3h18.05v6.017H6.975z"/><path fill="green" d="M12.992 10.95v6.017H6.975v6.017h6.017V29h6.017v-6.017h6.017v-6.016h-6.018V10.95Z"/>`),
		g.Group(children),
	)
}