package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSixCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="currentColor" d="M13.418 7.274a.5.5 0 0 0-.836-.548l.836.548Zm-4.361 4.831a.5.5 0 1 0 .836.548l-.836-.548ZM14.5 14a2.5 2.5 0 0 1-2.5 2.5v1a3.5 3.5 0 0 0 3.5-3.5h-1ZM12 16.5A2.5 2.5 0 0 1 9.5 14h-1a3.5 3.5 0 0 0 3.5 3.5v-1ZM9.5 14a2.5 2.5 0 0 1 2.5-2.5v-1A3.5 3.5 0 0 0 8.5 14h1Zm2.5-2.5a2.5 2.5 0 0 1 2.5 2.5h1a3.5 3.5 0 0 0-3.5-3.5v1Zm.582-4.774l-3.525 5.38l.836.547l3.525-5.379l-.836-.548Z"/></g>`),
		g.Group(children),
	)
}