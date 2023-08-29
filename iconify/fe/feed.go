package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Feed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feFeed0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFeed1" fill="currentColor"><path id="feFeed2" d="M6 19a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 2a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm12.971 0C18.473 12.408 11.592 5.527 3 5.029V3c9.941 0 18 8.059 18 18h-2.029Zm-5.012 0C13.478 15.17 8.829 10.522 3 10.041V8c7.18 0 13 5.82 13 13h-2.041Z"/></g></g>`),
		g.Group(children),
	)
}