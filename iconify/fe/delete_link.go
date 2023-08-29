package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteLink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feDeleteLink0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDeleteLink1" fill="currentColor" fill-rule="nonzero"><path id="feDeleteLink2" d="M11 9H7a3 3 0 0 0 0 6h4v2H7A5 5 0 0 1 7 7h4v2Zm2 6h4a3 3 0 0 0 0-6h-4V7h4a5 5 0 0 1 0 10h-4v-2Zm0-4h2a1 1 0 0 1 0 2h-2v-2Zm-4 0h2v2H9a1 1 0 0 1 0-2Z"/></g></g>`),
		g.Group(children),
	)
}