package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeModernizr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#d91b77" d="M2 22.72v-4.48h4.48v-4.48h4.48V9.28h4.48v13.44m1.12-13.44A13.44 13.44 0 0 1 30 22.72H16.56V9.28"/>`),
		g.Group(children),
	)
}