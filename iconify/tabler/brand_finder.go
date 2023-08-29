package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandFinder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 5a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1zm4 3v1m10-1v1"/><path d="M12.5 4c-.654 1.486-1.26 3.443-1.5 9h2.5c-.19 2.867.094 5.024.5 7"/><path d="M7 15.5c3.667 2 6.333 2 10 0"/></g>`),
		g.Group(children),
	)
}