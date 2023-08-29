package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkExternal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feLinkExternal0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feLinkExternal1" fill="currentColor"><path id="feLinkExternal2" d="M6 8h5v2H6v8h8v-5h2v5a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2Zm10.614-2H12V4h8v8h-2V7.442l-5.336 5.336l-1.414-1.414L16.614 6Z"/></g></g>`),
		g.Group(children),
	)
}