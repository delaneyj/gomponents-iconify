package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeCircleBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="m15.5 9l.172.172c1.333 1.333 2 2 2 2.828c0 .828-.667 1.495-2 2.828L15.5 15m-2.206-7.83L12 12l-1.294 4.83M8.5 9l-.172.172c-1.333 1.333-2 2-2 2.828c0 .828.667 1.495 2 2.828L8.5 15"/><path d="M7 3.338A9.954 9.954 0 0 1 12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12c0-1.821.487-3.53 1.338-5"/></g>`),
		g.Group(children),
	)
}