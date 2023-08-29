package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="fePot0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePot1" fill="currentColor"><path id="fePot2" d="M20 11v7a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-7H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2h-1ZM6 11v7h12v-7H6Zm5-5V5a1 1 0 0 1 2 0v1h6a1 1 0 0 1 0 2H5a1 1 0 1 1 0-2h6Z"/></g></g>`),
		g.Group(children),
	)
}