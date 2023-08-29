package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Three(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M25.87 14c2.322 0 4.13 1.831 4.13 4s-1.808 4-4.13 4h-4.088a2 2 0 1 0 0 4h4.087C28.192 26 30 27.831 30 30s-1.808 4-4.13 4h-4.088c-1.823 0-3.344-1.137-3.9-2.68a2 2 0 0 0-3.763 1.36c1.126 3.118 4.147 5.32 7.663 5.32h4.087C30.32 38 34 34.459 34 30a7.91 7.91 0 0 0-2.753-6A7.91 7.91 0 0 0 34 18c0-4.459-3.681-8-8.13-8h-4.088c-3.516 0-6.537 2.202-7.663 5.32a2 2 0 0 0 3.762 1.36c.557-1.543 2.078-2.68 3.901-2.68h4.087Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}