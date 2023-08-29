package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignOutAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.707 11.293l-4-4a1 1 0 0 0-1.414 1.414L12.586 11H2.05c-.03.33-.05.66-.05 1s.02.67.05 1h10.536l-2.293 2.293a1 1 0 1 0 1.414 1.414l4-4a1 1 0 0 0 0-1.414Z"/><path fill="currentColor" d="M12 2a10 10 0 0 0-9.95 9h10.536l-2.293-2.293a1 1 0 0 1 1.414-1.414l4 4a1 1 0 0 1 0 1.414l-4 4a1 1 0 0 1-1.414-1.414L12.586 13H2.05A10 10 0 1 0 12 2Z" opacity=".5"/>`),
		g.Group(children),
	)
}