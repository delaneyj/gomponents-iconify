package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FigmaMask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20"/><path d="M14 41.324C19.978 37.866 24 31.403 24 24c0-7.403-4.022-13.866-10-17.324C8.022 10.134 4 16.597 4 24c0 7.403 4.022 13.866 10 17.324Z"/></g>`),
		g.Group(children),
	)
}