package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.707 3.293a1 1 0 0 0-1.414 1.414l.795.795a6 6 0 0 0 .134 8.347l5.657 5.657a3 3 0 0 0 4.242 0l1.986-1.985l3.186 3.186a1 1 0 0 0 1.414-1.414l-3.885-3.885a.646.646 0 0 0-.017-.017L6.267 4.853a1 1 0 0 0-.03-.03l-1.53-1.53ZM12 4.758a6.002 6.002 0 0 1 7.778 9.091l-.136.136a1 1 0 0 1-1.414-1.414l.136-.136a4 4 0 0 0-5.657-5.657a1 1 0 0 1-1.414 0a4.003 4.003 0 0 0-.889-.671a1 1 0 1 1 .971-1.749c.215.12.423.253.625.4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}