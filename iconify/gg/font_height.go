package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M21 5V3H3v2h18Zm0 14v2H3v-2h18Z"/><path fill-rule="evenodd" d="M12 7.376a1 1 0 0 0-.967.576l-3.381 7.25a1 1 0 1 0 1.812.846L9.953 15h4.094l.489 1.048a1 1 0 1 0 1.813-.845l-3.381-7.251A1 1 0 0 0 12 7.376ZM13.115 13h-2.23L12 10.61L13.115 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}