package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrowserOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaBrowserOutline0"><g id="evaBrowserOutline1"><g id="evaBrowserOutline2" fill="currentColor"><path d="M18 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3Zm1 15a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-7h14ZM5 9V6a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v3Z"/><circle cx="8" cy="7.03" r="1"/><circle cx="12" cy="7.03" r="1"/></g></g></g>`),
		g.Group(children),
	)
}