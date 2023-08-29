package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feTag0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTag1" fill="currentColor" fill-rule="nonzero"><path id="feTag2" d="m4.981 14.887l4.132 4.132l9.906-9.906V4.981h-4.132L4.98 14.887ZM13.486 3.58a1.98 1.98 0 0 1 1.4-.58h4.133C20.113 3 21 3.887 21 4.981v4.132c0 .526-.209 1.03-.58 1.401l-9.906 9.906a1.981 1.981 0 0 1-2.802 0L3.58 16.288a1.981 1.981 0 0 1 0-2.802l9.906-9.906ZM16 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g></g>`),
		g.Group(children),
	)
}