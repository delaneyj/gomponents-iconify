package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerifiedUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M2 20v-1a7 7 0 0 1 7-7v0"/><path d="M15.804 12.313a1.618 1.618 0 0 1 2.392 0c.325.357.79.55 1.272.527a1.618 1.618 0 0 1 1.692 1.692c-.023.481.17.947.526 1.272c.705.642.705 1.75 0 2.392c-.356.325-.549.79-.526 1.272a1.618 1.618 0 0 1-1.692 1.692a1.618 1.618 0 0 0-1.272.526a1.618 1.618 0 0 1-2.392 0a1.618 1.618 0 0 0-1.272-.526a1.618 1.618 0 0 1-1.692-1.692a1.618 1.618 0 0 0-.527-1.272a1.618 1.618 0 0 1 0-2.392c.357-.325.55-.79.527-1.272a1.618 1.618 0 0 1 1.692-1.692c.481.023.947-.17 1.272-.527Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m15.364 17l1.09 1.09l2.182-2.18M9 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/></g>`),
		g.Group(children),
	)
}