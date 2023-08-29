package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feHeart0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feHeart1" fill="currentColor"><path id="feHeart2" d="M12 20c-2.205-.48-9-4.24-9-11a5 5 0 0 1 9-3a5 5 0 0 1 9 3c0 6.76-6.795 10.52-9 11Z"/></g></g>`),
		g.Group(children),
	)
}