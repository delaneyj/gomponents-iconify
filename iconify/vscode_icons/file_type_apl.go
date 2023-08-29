package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeApl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#d2d2d2" d="M30 28.275L16 2L2 28.275h12.162V30h3.676v-1.725Zm-12.162-3.449V13.161l6.215 11.665Zm-9.891 0l6.215-11.665v11.665Z"/>`),
		g.Group(children),
	)
}