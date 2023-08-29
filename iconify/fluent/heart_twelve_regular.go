package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.41 2.515a2.39 2.39 0 0 0-3.2.213c-.95.974-.946 2.558.008 3.536L5.75 9.887c.146.15.384.15.53 0l3.513-3.602a2.548 2.548 0 0 0-.01-3.535a2.396 2.396 0 0 0-3.45-.009l-.336.345l-.34-.35a2.498 2.498 0 0 0-.246-.22Zm1.638.924a1.396 1.396 0 0 1 2.018.009c.577.592.577 1.553.009 2.14l-.001.001l-3.06 3.138l-3.08-3.16a1.547 1.547 0 0 1-.008-2.141a1.394 1.394 0 0 1 2.014.009l.34.349a1 1 0 0 0 1.433 0l.335-.345Z"/>`),
		g.Group(children),
	)
}