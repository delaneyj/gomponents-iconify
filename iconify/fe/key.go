package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Key(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feKey0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feKey1" fill="currentColor" fill-rule="nonzero"><path id="feKey2" d="M11.9 11H22v4.004c0 .546-.45.996-1 .996s-1-.45-1-.996V13h-2v2c0 .55-.45 1-1 1s-1-.45-1-1v-2h-4.1A5.002 5.002 0 0 1 2 12a5 5 0 0 1 9.9-1ZM7 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g></g>`),
		g.Group(children),
	)
}