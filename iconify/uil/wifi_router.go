package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiRouter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.9 17a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-3 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm6 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1ZM15 8.5a1 1 0 0 1 1.73 0a1 1 0 0 0 1.36.37a1 1 0 0 0 .41-1.37a3 3 0 0 0-5.2 0a1 1 0 0 0 1.7 1Zm7-3a7 7 0 0 0-12.12 0a1 1 0 0 0 .37 1.37a1 1 0 0 0 .45.13a1 1 0 0 0 .87-.5a5 5 0 0 1 8.66 0a1 1 0 0 0 1.37.37A1 1 0 0 0 22 5.5ZM17.9 14h-1v-3a1 1 0 1 0-2 0v3h-10a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h13a3 3 0 0 0 3-3v-2a3 3 0 0 0-3-3Zm1 5a1 1 0 0 1-1 1h-13a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h13a1 1 0 0 1 1 1Z"/>`),
		g.Group(children),
	)
}