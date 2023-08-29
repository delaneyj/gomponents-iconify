package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VipLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.005 19h20v2h-20v-2Zm9-11h2v8h-2V8ZM7.97 8l-1.86 5.113L4.247 8H2.123l2.986 7.964h2L10.095 8H7.97Zm9.035 6v2h-2V8h4a3 3 0 0 1 0 6h-2Zm0-4v2h2a1 1 0 1 0 0-2h-2Zm-15-7h20v2h-20V3Z"/>`),
		g.Group(children),
	)
}