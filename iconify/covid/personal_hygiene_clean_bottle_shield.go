package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonalHygieneCleanBottleShield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M23 15.749a7.67 7.67 0 0 1-6 7.5a7.67 7.67 0 0 1-6-7.5V13.5a1.5 1.5 0 0 1 1.5-1.5h9a1.5 1.5 0 0 1 1.5 1.5v2.249Zm-5.999-.75v4.5m-2.251-2.25h4.501"/><path d="M8.5 20.253H4a3 3 0 0 1-3-3v-7.5a3 3 0 0 1 3-3h6a3 3 0 0 1 2.925 2.331M1 2.25l.622-.621A3 3 0 0 1 3.741.75h7.009M10 6.751H4v-1.5a1.5 1.5 0 0 1 1.5-1.5h3a1.5 1.5 0 0 1 1.5 1.5v1.5ZM6.999 3.75v-3"/></g>`),
		g.Group(children),
	)
}