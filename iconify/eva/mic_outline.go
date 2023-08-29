package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaMicOutline0"><g id="evaMicOutline1"><g id="evaMicOutline2" fill="currentColor"><path d="M12 15a4 4 0 0 0 4-4V6a4 4 0 0 0-8 0v5a4 4 0 0 0 4 4Zm-2-9a2 2 0 0 1 4 0v5a2 2 0 0 1-4 0Z"/><path d="M19 11a1 1 0 0 0-2 0a5 5 0 0 1-10 0a1 1 0 0 0-2 0a7 7 0 0 0 6 6.92V20H8.89a.89.89 0 0 0-.89.89v.22a.89.89 0 0 0 .89.89h6.22a.89.89 0 0 0 .89-.89v-.22a.89.89 0 0 0-.89-.89H13v-2.08A7 7 0 0 0 19 11Z"/></g></g></g>`),
		g.Group(children),
	)
}