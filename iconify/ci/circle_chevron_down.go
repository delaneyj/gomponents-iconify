package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleChevronDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm0-18a8 8 0 1 0 8 8a8.009 8.009 0 0 0-8-8Zm-.05 11.5L7 10.55l1.414-1.414l3.536 3.535l3.536-3.535L16.9 10.55l-4.95 4.95Z"/>`),
		g.Group(children),
	)
}