package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Music(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.954 12.304V1a1 1 0 0 1 1-1h.028a1 1 0 0 1 1 1v14.065c.197 1.969-1.42 3.99-3.874 4.693c-2.69.772-5.368-.333-5.98-2.468c-.612-2.135 1.073-4.491 3.764-5.263c1.47-.421 2.935-.283 4.062.277zm-2.4 5.521c1.698-.487 2.645-1.81 2.37-2.77c-.276-.961-1.78-1.582-3.478-1.095c-1.698.487-2.645 1.81-2.37 2.771c.276.96 1.78 1.581 3.478 1.094z"/>`),
		g.Group(children),
	)
}