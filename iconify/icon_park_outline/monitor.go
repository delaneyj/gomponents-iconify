package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M4 10h32v28H4z"/><path stroke-linecap="round" d="m44 14l-8 6.75v6.5L44 34V14Z" clip-rule="evenodd"/><path stroke-linecap="round" d="m17 19l6 5l-6 5"/></g>`),
		g.Group(children),
	)
}