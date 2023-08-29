package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feGift0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feGift1" fill="currentColor"><path id="feGift2" d="M19 12v8a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-8a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h1.17A3 3 0 1 1 12 5a3 3 0 1 1 5.83 1H19a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2Zm-8-4H5v2h6V8Zm2 0v2h6V8h-6Zm-6 4v8h10v-8H7Zm2-6h1V5a1 1 0 1 0-1 1Zm6 0a1 1 0 1 0-1-1v1h1Z"/></g></g>`),
		g.Group(children),
	)
}