package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feTerminal0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTerminal1" fill="currentColor" fill-rule="nonzero"><path id="feTerminal2" d="M4 4h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Zm0 2v12h16V6H4Zm8 8h6v2h-6v-2Zm-1.015-2.429L7.45 15.107l-1.414-1.415l2.12-2.12l-2.12-2.122L7.45 8.036l3.535 3.535Z"/></g></g>`),
		g.Group(children),
	)
}