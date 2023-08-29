package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M11.473 9a4.5 4.5 0 0 0-8.72-.99A3 3 0 0 0 3 14h8.5a2.5 2.5 0 1 0-.027-5z"/><path d="M14.544 9.772a3.506 3.506 0 0 0-2.225-1.676a5.502 5.502 0 0 0-6.337-4.002a4.002 4.002 0 0 1 7.392.91a2.5 2.5 0 0 1 1.17 4.769z"/></g>`),
		g.Group(children),
	)
}