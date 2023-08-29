package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElipsisH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feElipsisH0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feElipsisH1" fill="currentColor"><path id="feElipsisH2" d="M18 14a2 2 0 1 1 0-4a2 2 0 0 1 0 4ZM6 14a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm6 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/></g></g>`),
		g.Group(children),
	)
}