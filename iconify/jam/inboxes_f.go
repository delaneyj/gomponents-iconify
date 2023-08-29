package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxesF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 6v5H0V6h6.126a4.002 4.002 0 0 0 7.748 0H20zm0-2H0l5-4h10l5 4zm0 9v5a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-5h6.126a4.002 4.002 0 0 0 7.748 0H20zM8.268 13h3.464a2 2 0 0 1-3.464 0zm-.002-6.988h3.464a2 2 0 0 1-3.464 0z"/>`),
		g.Group(children),
	)
}