package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuddiconsTopics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.44 1.66c-.59-.58-1.54-.58-2.12 0L2.66 7.32c-.58.58-.58 1.53 0 2.12c.6.6 1.56.56 2.12 0l5.66-5.66a1.5 1.5 0 0 0 0-2.12zm2.83 2.83a1.49 1.49 0 0 0-2.12 0l-5.66 5.66a1.49 1.49 0 0 0 0 2.12c.6.6 1.56.55 2.12 0l5.66-5.66c.58-.58.58-1.53 0-2.12zm1.06 6.72l4.18 4.18c.59.58.59 1.53 0 2.12s-1.54.59-2.12 0l-4.18-4.18l-1.77 1.77c-.59.58-1.54.58-2.12 0c-.59-.59-.59-1.54 0-2.13l5.66-5.65a1.49 1.49 0 0 1 2.12 0c.58.58.58 1.53 0 2.12zM5 15c0-1.59-1.66-4-1.66-4S2 13.78 2 15s.6 2 1.34 2h.32C4.4 17 5 16.59 5 15z"/>`),
		g.Group(children),
	)
}