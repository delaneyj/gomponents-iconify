package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopMinimalisticLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 9c0-2.828 0-4.243.879-5.121C4.757 3 6.172 3 9 3h6c2.828 0 4.243 0 5.121.879C21 4.757 21 6.172 21 9v5c0 1.886 0 2.828-.586 3.414C19.828 18 18.886 18 17 18H7c-1.886 0-2.828 0-3.414-.586C3 16.828 3 15.886 3 14V9Z"/><path stroke-linecap="round" d="M22 21H2m13-6H9"/></g>`),
		g.Group(children),
	)
}