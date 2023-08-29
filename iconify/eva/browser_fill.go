package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrowserFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaBrowserFill0"><g id="evaBrowserFill1"><path id="evaBrowserFill2" fill="currentColor" d="M18 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3Zm-6 3a1 1 0 1 1-1 1a1 1 0 0 1 1-1ZM8 6a1 1 0 1 1-1 1a1 1 0 0 1 1-1Zm11 12a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-7h14Z"/></g></g>`),
		g.Group(children),
	)
}