package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a3 3 0 1 0 0 6a3 3 0 0 0 0-6ZM7 6a5 5 0 1 1 6 4.9V17h-2v-6.1A5.002 5.002 0 0 1 7 6Zm-3.895 5H8v2H4.895l-.778 7h15.766l-.778-7H16v-2h4.895l1.222 11H1.883l1.222-11Z"/>`),
		g.Group(children),
	)
}