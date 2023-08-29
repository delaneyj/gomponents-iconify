package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeAlDal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#9F4246" d="M11.6 8c-.2-.6-.8-1-1.5-1H10c-.5 0-1.1.4-1.4 1L2 25h3.8l1-2.8h6.6l1 2.8h3.8L11.6 8zM7.9 19l2.2-6.1l2.2 6.1H7.9zm15.9 2.7V7.1h-3.5v15.6c0 1.2 1 2.2 2.3 2.3H30v-3.3h-6.2z"/>`),
		g.Group(children),
	)
}