package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationPinOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.685 4.627a6.78 6.78 0 0 1 7.108 1.58a6.78 6.78 0 0 1 1.58 7.108a1 1 0 1 0 1.88.681a8.78 8.78 0 0 0-2.046-9.203a8.78 8.78 0 0 0-9.204-2.046a1 1 0 1 0 .682 1.88Zm-4.982 1.49l-1.41-1.41a1 1 0 0 1 1.414-1.414L6.72 5.306a.978.978 0 0 1 .02.02l10.96 10.96l.007.007l.007.006l2.993 2.994a1 1 0 0 1-1.414 1.414L17 18.414l-4.293 4.293a1 1 0 0 1-1.414 0l-5.5-5.5a8.78 8.78 0 0 1-1.09-11.09Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}