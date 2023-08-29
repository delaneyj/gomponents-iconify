package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiController(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 30h20a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2Zm0-8h20v6H6Z"/><circle cx="9" cy="25" r="1" fill="currentColor"/><circle cx="16" cy="14.5" r="1.5" fill="currentColor"/><path fill="currentColor" d="M10.783 10.332a7 7 0 0 1 10.434 0l-1.49 1.334a5 5 0 0 0-7.453 0Z"/><path fill="currentColor" d="M7.2 7.4a11.002 11.002 0 0 1 17.6 0l-1.6 1.2a9 9 0 0 0-14.401 0Z"/>`),
		g.Group(children),
	)
}